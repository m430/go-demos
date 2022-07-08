package chapter42

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	compareTests := map[string]struct {
		a, b string
		i    int
	}{
		"compareTowEmptyString":     {"", "", 0},
		"compareSecondParamIsEmpty": {"a", "", 1},
		"compareFirstParamIsEmpty":  {"", "a", -1},
	}

	for name, tt := range compareTests {
		t.Run(name, func(t *testing.T) {
			cmp := strings.Compare(tt.a, tt.b)
			if cmp != tt.i {
				t.Errorf("[%s] want %v, but compare(%q, %q) = %v", name, tt.i, tt.a, tt.b, cmp)
			}
		})
	}
}

// 总结：
// 1. 表驱动测试简单和紧凑
// 2. 数据即测试
// 3. 结合子测试，可单独运行某个数据项的测试
