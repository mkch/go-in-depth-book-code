package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	_ "embed"
)

//go:embed upload.html
var uploadPage []byte

func main() {
	http.HandleFunc("GET /upload", handleUpload)       // 显示上传页面
	http.HandleFunc("POST /send_file", handleSendFile) // 上传文件
	http.HandleFunc("GET /download", handleDownload)   // 下载文件

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadPage)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	task := takeTask()
	if task == nil {
		// 没有上传方
		http.Error(w, "No task", http.StatusBadRequest)
		return
	}
	select {
	case reader, ok := <-task.Content:
		if !ok {
			// 上传方已取消
			http.Error(w, "Cancelled", http.StatusNotFound)
			return
		}
		// 开始下载
		setHeaders(w, task)
		_, err := io.Copy(w, reader)
		task.DownloadErr <- err
	case <-r.Context().Done():
		// 客户端已取消下载
	}
}

func handleSendFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")
	fileSize, err := strconv.ParseUint(r.URL.Query().Get("size"), 10, 64)
	if err != nil || fileName == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	task := newTask(fileName, fileSize)
	if !setTask(task) {
		http.Error(w, "Task pending", http.StatusBadRequest)
		return
	}

	// 设置等待超时
	timeout, cancel := context.WithTimeout(context.Background(),
		time.Minute*5)
	defer cancel()

	select {
	case <-r.Context().Done():
		// 客户端已取消上传
		close(task.Content)
		takeTask() // 丢弃当前任务
	case <-timeout.Done():
		// 等待已超时
		close(task.Content)
		takeTask() // 丢弃当前任务
		http.Error(w, "Timed out", http.StatusBadRequest)
	case task.Content <- r.Body:
		// 下载方已开始下载
		if <-task.DownloadErr != nil {
			// 一旦开始读 body, 就必须读完, 否则前端的 fetch() 不会返回
			io.Copy(io.Discard, r.Body)
			http.Error(w, "Error downloading", http.StatusBadRequest)
		}
		return
	}
}

// setHeaders 为 w 设置必要的 header 以便客户端能确定文件大小并
// 以正确的文件名下载文件.
func setHeaders(w http.ResponseWriter, task *task) {
	w.Header().Set("Content-Length",
		strconv.FormatUint(task.FileSize, 10))
	w.Header().Set("Content-Disposition",
		fmt.Sprintf(`attachment; filename*=utf-8''%v`,
			url.PathEscape(task.FileName)))
	w.Header().Set("Content-Type", "application/octet-stream")
}

type task struct {
	FileName    string         // 文件名
	FileSize    uint64         // 文件大小
	Content     chan io.Reader // 上传方向下载方传递文件内容
	DownloadErr chan error     // 下载方向上传方传递错误
}

func newTask(name string, size uint64) *task {
	return &task{name, size, make(chan io.Reader), make(chan error)}
}

var curTaskLock sync.Mutex
var curTask *task

// setTask 设置当前任务为 task
// 如果当前有待下载任务, 返回 false.
func setTask(task *task) bool {
	curTaskLock.Lock()
	defer curTaskLock.Unlock()
	if curTask != nil {
		return false
	}
	curTask = task
	return true
}

// taskTask 取出当前任务, 并设置当前无任务.
// 如果当前无任务, 返回 nil.
func takeTask() (t *task) {
	curTaskLock.Lock()
	defer curTaskLock.Unlock()
	t = curTask
	curTask = nil
	return
}
