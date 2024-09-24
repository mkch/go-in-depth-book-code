package escape

func f(*int) {}

var fvar func(*int) = f

func F(a int) *int {
	var n1 int
	var ary []*int
	ary[0] = &n1

	var n2 int
	var s []*int = make([]*int, 1)
	s[0] = &n2

	var (
		n3 int
		n4 int
	)
	var m map[*int]*int
	m[&n3] = &n4

	var n5 int
	f(&n5)
	var n6 int
	fvar(&n6)

	var p = new(int)
	*p = 1

	var b int
	var c = a + b
	return &c
}

// 在此目录运行 go build -gcflags="-l -m" 查看结果
