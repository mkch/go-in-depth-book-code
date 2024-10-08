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
	http.HandleFunc("GET /upload", handleUpload)              // 显示上传页面
	http.HandleFunc("POST /send_file/{code}", handleSendFile) // 上传文件
	http.HandleFunc("GET /download/{code}", handleDownload)   // 下载文件

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadPage)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	task := getTask(r.PathValue("code"))
	if task == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	select {
	case content, ok := <-task.Content:
		if !ok {
			// 上传方已取消
			http.Error(w, "Cancelled", http.StatusNotFound)
			return
		}
		// 开始下载
		setHeaders(w, task)
		_, err := io.Copy(w, content)
		task.DownloadErr <- err
	case <-r.Context().Done():
		// 客户端已取消下载
	}
}

func handleSendFile(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	fileName := r.URL.Query().Get("name")
	fileSize, err := strconv.ParseUint(r.URL.Query().Get("size"), 10, 64)
	if err != nil || code == "" || fileName == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// 设置等待超时
	timeout, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	task := addTask(code, fileName, fileSize)
	if task == nil {
		http.Error(w, "Task already exists", http.StatusBadRequest)
		return
	}

	select {
	case <-r.Context().Done():
		// 客户端已取消上传
	case <-timeout.Done():
		// 等待已超时
		http.Error(w, "Timed out", http.StatusBadRequest)
	case task.Content <- r.Body:
		// 下载方已开始下载
		if <-task.DownloadErr != nil {
			// 一旦开始读 body, 就必须读完, 否则前端的 fetch() 不会返回
			io.Copy(io.Discard, r.Body)
			http.Error(w, "Failed to download", http.StatusBadRequest)
		}
	}
	close(task.Content)
	removeTask(code)
}

// setHeaders 为 w 设置必要的 header 以便客户端能确定文件大小并
// 以正确的文件名下载文件.
func setHeaders(w http.ResponseWriter, task *task) {
	w.Header().Set("Content-Length", strconv.FormatUint(task.FileSize, 10))
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	w.Header().Set("Content-Disposition",
		fmt.Sprintf(`attachment; filename*=utf-8''%v`, url.PathEscape(task.FileName)))
	w.Header().Set("Content-Type", "application/octet-stream")
}

// addTask 添加一个传输任务. 如果 code 已存在则返回 nil.
func addTask(code string, fileName string, fileSize uint64) *task {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	if tasks[code] != nil {
		return nil // 任务已存在
	}
	task := &task{
		FileName:    fileName,
		FileSize:    fileSize,
		Content:     make(chan io.Reader),
		DownloadErr: make(chan error),
	}
	tasks[code] = task
	return task
}

// getTask 从任务列表中取出 code 对应的等待下载的 task,
// 并设置该 task 为正在下载.
// 如果不存在则返回 nil.
func getTask(code string) *task {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	t := tasks[code]
	if t == nil || t.Downloading {
		return nil
	}
	t.Downloading = true
	return t
}

// removeTask 从任务列表移除 code 对应的 task.
func removeTask(code string) {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	delete(tasks, code)
}

type task struct {
	FileName    string         // 文件名
	FileSize    uint64         // 文件大小
	Content     chan io.Reader // 上传方向下载方传递文件内容
	DownloadErr chan error     // 下载方向上传方传递错误
	Downloading bool           // 是否正在下载
}

var tasksLock sync.RWMutex
var tasks = make(map[string]*task)
