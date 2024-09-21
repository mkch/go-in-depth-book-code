package bench

import "testing"

func TestLastDot(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a.b.c", args{"a.b.c"}, 3},
		{"a.", args{"a."}, 1},
		{".", args{"."}, 0},
		{"EMPTY", args{""}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDot(tt.args.str); got != tt.want {
				t.Errorf("LastDot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLastDot(b *testing.B) {
	for range b.N {
		LastDot("abc.ef")
	}
}
