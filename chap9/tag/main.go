package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type S struct {
	F1 int "Some tag"
	F2 int `k1:"v1" k2:"v2"`
}

func main() {
	var a S
	s := reflect.TypeOf(a)
	f1, f2 := s.Field(0), s.Field(1)
	// 原样输出字段的 tag
	fmt.Println(f1.Tag)
	fmt.Println(f2.Tag)
	// 如果 tag 为 key:"value" 对
	// 还可以取出 key 对应的 value
	fmt.Println(f2.Tag.Get("k1"))

	decodeRepo()
}

type Repo struct {
	MainBranch string `json:"master_branch"`
}

func decodeRepo() {
	data := []byte(`{"master_branch": "branch1"}`)
	var repo Repo
	json.Unmarshal(data, &repo)
	fmt.Println(repo)
}
