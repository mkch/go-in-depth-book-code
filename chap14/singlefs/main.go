package singlefs

import (
	"io"
	"io/fs"
	"os"
)

// SingleFileFS 是一个只包含一个文件的 fs.FS
type SingleFileFS struct {
	fs.FS
	file fs.FileInfo // 唯一的文件
}

// New 创建一个只包含一个文件的文件系统.
func New(dir, file string) (*SingleFileFS, error) {
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
	// 委托给 fs.FS
	f, err := s.FS.Open(name)
	if err != nil {
		return nil, err
	}
	// 打开目录自身
	if name == "." {
		return &SingleFileDir{ReadDirFile: f.(fs.ReadDirFile), file: s.file}, nil
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}
	// 打开唯一的文件
	if os.SameFile(s.file, info) {
		return f, nil
	}
	// 试图打开其他文件, 返回 "该文件不存在"
	f.Close()
	return nil, os.ErrNotExist

}

// SingleFileDir 是一个只包含一个文件的 fs.ReadDirFile
type SingleFileDir struct {
	fs.ReadDirFile
	file fs.FileInfo // 如果为 nil,表示已经遍历完毕
}

func (dir *SingleFileDir) ReadDir(n int) (entries []fs.DirEntry, err error) {
	if dir.file == nil {
		if n > 0 {
			err = io.EOF
		}
		return
	}
	entries = []fs.DirEntry{fs.FileInfoToDirEntry(dir.file)}
	dir.file = nil // 遍历完毕
	return
}
