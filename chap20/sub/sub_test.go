package sub

import "testing"

func TestFoo(t *testing.T) {
	// <公共初始化代码>
	t.Run("A", func(t *testing.T) {
		// t.Run("AA", func(t *testing.T) { t.Log("AA") })
		// t.Run("AA", func(t *testing.T) { t.Log("AA#01") })
		t.Log("A")
	})
	t.Run("A", func(t *testing.T) { t.Log("A#01") })
	t.Run("B", func(t *testing.T) { t.Log("B") })
	// <公共清理代码>
}
