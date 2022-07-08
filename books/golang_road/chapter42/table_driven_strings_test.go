package chapter42

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	compareTests := []struct {
		a, b string
		i    int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "a", -1},
	}

	for _, tt := range compareTests {
		cmp := strings.Compare(tt.a, tt.b)
		if cmp != tt.i {
			t.Errorf("want %v, but compare(%q, %q) = %v", tt.i, tt.a, tt.b, cmp)
		}
	}
}

// 总结：
// 1. 表驱动测试简单和紧凑
// 2. 数据即测试
// 3. 结合子测试，可单独运行某个数据项的测试
