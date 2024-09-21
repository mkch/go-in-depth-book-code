package main

type T struct {
	a int
}

func (tv T) Mv(a int) int          { return a + 1 } // 接收者为 T
func (tp *T) Mp(f float32) float32 { return 1.0 }   // 接收者为 *T

var t T

func main() {
	T.Mv(t, 1)       // 相当于 t.Mv(1)
	(*T).Mp(&t, 1.0) // 相当于 t.Mp(1.0)
	(*T).Mv(&t, 1)   // 相当于 t.Mv(1)
	// T.Mp(&t, 1.0)
}
