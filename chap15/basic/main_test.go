package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if r := add(1, 2); r != 3 {
		// 如果被测函数的行为不符合预期
		// 则报告测试失败
		t.Errorf("add(1, 2) = %v, want %v", r, 3)
	}
}

func decode() []byte      { return nil }
func encode([]byte) error { return nil }

func TestDecodeEncode(t *testing.T) {
	data := decode() // 假设 decode 不应该返回 nil
	if data == nil {
		// 使用 Fatal() 终止当前测试
		// 因为后续 encode() 的测试依赖 data
		t.Fatal("decode() returns nil")
	}
	// 使用 decode() 返回的 data 来测试 encode()
	if err := encode(data); err != nil {
		t.Fatalf("encode() returns %v, want nil", err)
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"positive", args{1, 2}, 3},
		{"negative", args{-1, -2}, -3},
		{"zero", args{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
