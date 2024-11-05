package res

// Resource 描述特定类型的资源.
type Resource struct {
	// ...
}

// globalRes 是一个全局共享资源.
var globalRes Resource

// Res 返回全局共享资源的引用.
func Res() *Resource {
	return &globalRes
}
