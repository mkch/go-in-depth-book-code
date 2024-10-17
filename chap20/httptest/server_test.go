package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleConcat(t *testing.T) {
	// 模拟一个 Request
	r := httptest.NewRequest("GET", "/?a=abc&b=def", nil)
	// 模拟一个 ResponseWriter
	w := httptest.NewRecorder()
	// 调用被测函数
	handleConcat(w, r)
	response := w.Result()
	// 测试状态码
	if response.StatusCode != 200 {
		t.Fatalf("wrong status: %v, want %v", response.StatusCode, 200)
	}
	// 测试回应数据
	if data, err := io.ReadAll(response.Body); err != nil {
		t.Fatal(err)
	} else if str := string(data); str != "abcdef" {
		t.Fatalf("wrong response: %v, want %v", str, "abcdef")
	}
}

func TestMustHttp2InHttp2(t *testing.T) {
	// 创建测试 server 但不启动
	server := httptest.NewUnstartedServer(http.HandlerFunc(mustHttp2))
	// 启用 HTTP2
	server.EnableHTTP2 = true
	// 启动 TLS
	server.StartTLS()
	defer server.Close()
	// 模拟访问
	response, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("wrong status: %v, want %v", response.StatusCode, 200)
	}
}

func TestMustHttp2InHttp1(t *testing.T) {
	// 创建一个非 HTTP2 的测试服务器
	server := httptest.NewServer(http.HandlerFunc(mustHttp2))
	defer server.Close()
	// 模拟访问
	response, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 400 {
		t.Fatalf("wrong status: %v, want %v", response.StatusCode, 400)
	}
}
