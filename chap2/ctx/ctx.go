package ctx

import "context"

type User struct{}

// key 是定义在本包中的私有类型,
// 用此类型来定义关键字可避免与其他包中定义的关键字冲突.
type key int

// userKey 是 Context 中 User 的关键字
var userKey key

func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u)
}
