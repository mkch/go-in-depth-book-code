// Package demo 演示了 Go doc comment 的使用方法。
//
// 关于 doc comment 可参看 [go/doc] 和 [go/doc/comment]。
package demo

import (
	"errors"
	"math"

	"golang.org/x/net/html"
)

// Point 表示二维平面上的一个点的坐标。
type Point struct {
	X int // 横坐标值
	Y int // 纵坐标值
}

// Distance 计算两个点 pt1 和 pt2 之间的距离。
func Distance(pt1, pt2 *Point) int {
	var nX, nY = pt1.X - pt2.X, pt1.Y - pt2.Y
	return int(math.Sqrt(float64(nX*nX + nY*nY)))
}

// 一周的 7 天
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	MaxDigit = 9 // 一位数的最大值
	MinDigit = 1 // 一位数的最小值
)

// Version 是此程序库当前的版本。
//
// 每次升级本库时都会同步修改此版本号。
const Version = "1.2.3"

// Token 代表编译器词法分析阶段生成的标记。
type Token uint16

const (
	ID      Token = iota + 1 // 标识符
	PLUS                     // '+'
	MINUS                    // '-'
	INTEGER                  // 整数
)

var (
	Error1 error = errors.New("error1") // 函数 F 在 ... 时会返回此错误
	Error2       = errors.New("error2") // 函数 F2 在 ... 时会返回此错误
)

// Ops 是运算符 Token 和运算符字符串的对应表。
var Ops = map[Token]string{
	MINUS: "-",
	PLUS:  "+",
}

/*
下面是一个“段落”：

段落第一句话。
段落第二句话，比较长。段落第三句话。
*/
var ParagraphDemo int

/*
下面为两个标题：

# 第一

这是第一个标题下的内容

# 第二

这是第二个标题下的内容

#这不是一个标题，因为缺少空格

# 这也不是一个标题，
# 因为它有多行。

# 这也不是一个标题，
它同样有多行

在一个段落中，没有空行分割，
# 这也不是一个标题

	#这个不是标题，因为它有缩进
*/
var HeadingDemo int

/*
这是一个链接[链接1]。这是[链接2]。

这是一个错误的链接[链接4]

[链接1]: http://www.example.com/1
[链接2]: http://www.example.com/2

[链接3]: http://www.example.com/3
*/
var LinkDemo int

var _ html.NodeType

/*
这里引用本包内的类型 [Point]。引用本包内的函数 [Distance]。

引用其他包 [io],以及其他包中的标识符 [io.Reader] 和 [*bytes.Buffer]。

使用导入路径 [golang.org/x/mod/modfile.Comment]。

使用本文件中导入的包名 [html.Node]。
*/
var DocLinkDemo int

/*
这是第一个列表：
  - A
  - B
  - C

这是第二个列表：
  - 第一项 内容可以
    续行
  - 第二项

这是第三个列表
  - X
  - Y
  - Z

这是一个数字列表
 1. X
 2. Y
 3. Z
*/
var ListDemo int

/*
下面为一段 Go 代码：

	var a = 0

	func Distance(pt1, pt2 *Point) int {
		var nX, nY = pt1.X - pt2.X, pt1.Y - pt2.Y
		return int(math.Sqrt(float64(nX*nX + nY*nY)))
	}

下面为预格式化的文字：

	     12
	  9      3
	4   5   2   1
*/
var CodeBlockDemo int
