package escape

func f(*int) {}

var fvar func(*int) = f

func F(a int) *int {
	var n1 int
	var ary []*int
	ary[0] = &n1

	var n2 int                     // 第 12 行
	var s []*int = make([]*int, 1) // 第 13 行
	s[0] = &n2

	var (
		n3 int // 第 17 行
		n4 int // 第 18 行
	)
	var m map[*int]*int
	m[&n3] = &n4

	var n5 int
	f(&n5)
	var n6 int // 第 25 行
	fvar(&n6)

	var p = new(int) // 第 28 行
	*p = 1

	var b int
	var c = a + b // 第 32 行
	return &c
}

// 在此目录运行 go build -gcflags="-l -m" 查看结果
