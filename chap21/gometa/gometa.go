package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://github.com/mkch/iter2?go-get=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	meta := regexp.MustCompile(`\<meta\sname\=\"go\-import\".+\>`).
		Find(body)
	fmt.Println(string(meta))
}
