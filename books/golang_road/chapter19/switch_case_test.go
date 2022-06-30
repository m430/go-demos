package chapter19

import (
	"fmt"
	"testing"
)

func TestSwitchCase(t *testing.T) {
	var n = 1

	switch n {
	case 1, 3, 5, 7:
		fmt.Println("odd")
	case 2, 4, 6, 8:
		fmt.Println("even")
	default:
		fmt.Println("unknown")
	}
}

// 总结1：GO语言中case默认都是自动中断的，这一点和C语言不一样，省略了自己break的烦恼
