package hashmap_test

import (
	"fmt"
	"hashmap"
	"strconv"
	"testing"
)

func TestHashMap(t *testing.T) {
	m := hashmap.New[string, int](stringHash)
	if l := m.Len(); l != 0 {
		t.Fatal(l)
	}
	m.Set("0", 0)
	if v, ok := m.Get("0"); !ok || v != 0 {
		t.Fatal(v, ok)
	}

	for i := range 999 {
		m.Set(strconv.Itoa(i), i)
	}
	for i := range 999 {
		if v, ok := m.Get(strconv.Itoa(i)); !ok || v != i {
			t.Fatal(v, ok)
		}
	}
}

// stringHash 取
func stringHash(str string) uintptr {
	var h uintptr
	for _, r := range str {
		h = uintptr(r) + 31*h
	}
	return h
}

func ExampleHashMap() {
	// stringHash 是一个取 string 哈希值的函数
	stringHash := func(str string) uintptr {
		var h uintptr
		for _, r := range str {
			h = uintptr(r) + 31*h
		}
		return h
	}
	// 创建 HashMap
	m := hashmap.New[string, int](stringHash)
	// 写入
	m.Set("Red", 0xFF0000)
	m.Set("Green", 0x00FF00)
	m.Set("Blue", 0x0000FF)
	// 读取
	r, ok := m.Get("Red")
	fmt.Printf("%#x %v\n", r, ok)
	// 删除
	m.Delete("Red")
	// 遍历
	m.Range(func(kv hashmap.KeyValue[string, int]) bool {
		fmt.Printf("%v %#x\n", kv.Key, kv.Value)
		return true
	})
}
