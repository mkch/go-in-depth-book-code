package singlefs_test

import (
	"io"
	"io/fs"
	"os"
	"singlefs"
	"testing"
)

func TestSingleFileFSNotExists(t *testing.T) {
	fs, err := singlefs.New("testdata", "not exists")
	if fs != nil {
		t.Fatal(fs)
	}
	if err == nil {
		t.Fatal(err)
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
}

func TestSingleFileFS(t *testing.T) {
	fsys, err := singlefs.New("testdata", "b")
	if err != nil {
		t.Fatal(err)
	}
	f, err := fsys.Open("b")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if info.Name() != "b" {
		t.Fatal(info.Name())
	}

	dir, err := fsys.Open("..")
	if dir != nil {
		t.Fatal(dir)
	}
	if err == nil {
		t.Fatal(err)
	}

	dir, err = fsys.Open(".")
	if err != nil {
		t.Fatal(err)
	}
	defer dir.Close()
	df := dir.(fs.ReadDirFile)
	entries, err := df.ReadDir(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatal(entries)
	} else if name := entries[0].Name(); name != "b" {
		t.Fatal(name)
	}
	entries, err = df.ReadDir(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Fatal(entries)
	}
	entries, err = df.ReadDir(1)
	if err != io.EOF {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Fatal(entries)
	}
}
