package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Millisecond*500)
	defer cancel()

	r, err := Search(ctx, "golang")

	if err != nil {
		log.Println(err)
	} else {
		defer r.Body.Close()
		fmt.Fprintf(os.Stdout, "%v\n%v\n", r.Source, r.Status)
		io.Copy(os.Stdout, r.Body)
	}
}

func Search(ctx context.Context, keyword string) (*Result, error) {
	bingURL := bingSearchURL(keyword)
	sogouURL := sogouSearchURL(keyword)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bingRequest, err := http.NewRequestWithContext( // 1
		ctx, http.MethodGet, bingURL.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	sogouRequest, err := http.NewRequestWithContext( // 2
		ctx, http.MethodGet, sogouURL.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	ch := make(chan *result)

	go func() {
		ch <- NewResult(http.DefaultClient.Do(bingRequest))
	}()

	go func() {
		ch <- NewResult(http.DefaultClient.Do(sogouRequest))

	}()

	// 等待并发搜索结果
	r := <-ch
	// 较快的一个搜索已经返回，取消 Context
	cancel()

	// 清理落后者
	failure := <-ch
	if failure.Body != nil {
		io.Copy(io.Discard, failure.Body)
		failure.Body.Close()
	}

	if r.Err != nil {
		return nil, r.Err
	}
	return &r.Result, nil
}

// bingSearchURL 返回一个 cn.bing.com 的搜索 URL.
func bingSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://cn.bing.com/search")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("q", keyword)
	u.RawQuery = q.Encode()
	return u
}

// sogouSearchURL 返回一个 sogou.com 的搜索 URL.
func sogouSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://sogou.com/web")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("query", keyword)
	u.RawQuery = q.Encode()
	return u
}

// Result 是搜索结果.
type Result struct {
	Source string // 来源
	Status string
	Body   io.ReadCloser
}

type result struct {
	Result
	Err error
}

func NewResult(r *http.Response, err error) *result {
	if err != nil {
		return &result{Err: err}
	}
	return &result{
		Result: Result{
			Source: r.Request.URL.Host,
			Status: r.Status,
			Body:   r.Body},
	}
}
