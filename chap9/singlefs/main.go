package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func main() {
	fsys, err := NewSingleFileFS(".", "main.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(fsys.Open("main.GO"))
	fmt.Println(fsys.Open("go.mod"))
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
}

// SingleFileFS 是一个只包含一个文件的 fs.FS
type SingleFileFS struct {
	fs.FS
	file fs.FileInfo // 唯一的文件
}

// NewSingleFileFS 创建一个只包含一个文件的文件系统.
func NewSingleFileFS(dir, file string) (*SingleFileFS, error) {
	fsys := os.DirFS(dir)
	f, err := fsys.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return &SingleFileFS{fsys, fileInfo}, nil
}

func (s *SingleFileFS) Open(name string) (fs.File, error) {
	// 1. 委托给 fs.FS
	f, err := s.FS.Open(name)
	if err != nil {
		return nil, err
	}

	if name == "." { // 打开目录自身
		return &SingleFileDir{ReadDirFile: f.(fs.ReadDirFile), file: s.file}, nil
	} else {
		info, err := f.Stat()
		if err != nil {
			return nil, err
		}
		if os.SameFile(s.file, info) { // 打开唯一的文件
			return f, nil
		}
	}
	return nil, os.ErrNotExist // 试图打开其他文件,返回"该文件不存在"

}

// SingleFileDir 是一个只包含一个文件的 fs.ReadDirFile
type SingleFileDir struct {
	fs.ReadDirFile
	file fs.FileInfo // 如果为 nil,表示已经遍历完毕
}

func (dir *SingleFileDir) ReadDir(n int) ([]fs.DirEntry, error) {
	if dir.file == nil {
		if n <= 0 {
			return nil, nil
		}
		return nil, io.EOF
	}
	defer func() { dir.file = nil }() // 遍历完毕
	return []fs.DirEntry{fs.FileInfoToDirEntry(dir.file)}, nil
}
