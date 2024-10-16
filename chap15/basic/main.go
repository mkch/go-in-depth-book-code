package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func main() {
	Wrap2()
	PathError()
}

// ParseUserID 解析一个用户ID, 如果解析失败 err 将为非 nil
func ParseUserID(s string) (id int, err error) {
	id, err = strconv.Atoi(s)
	if err != nil {
		err = fmt.Errorf("parsing user id %v: %w", id, err)
	}
	return
}

func Wrap2() {
	var err1 = errors.New("reason1")
	var err2 = errors.New("reason2")
	var err = fmt.Errorf("failed: %w and %w", err1, err2)
	fmt.Println(err.Error())

	type Unwrapper interface {
		Unwrap() []error
	}

	wrapped := err.(Unwrapper).Unwrap()
	fmt.Println(wrapped)
}

func PathError() {
	f, err := os.Open("file1")
	if err != nil {
		defer f.Close()
	}
	if errors.Is(err, fs.ErrPermission) {
		fmt.Println("permission denied")
	}
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Println(pathErr)
	}
}
