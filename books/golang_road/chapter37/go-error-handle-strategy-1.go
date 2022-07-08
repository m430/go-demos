package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

// errors.Is会沿着该包装错误所在的错误链进行比较。直到匹配错误
func main() {
	err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err2: %w", err1)

	if errors.Is(err2, ErrSentinel) {
		println("err is ErrSentinel")
		return
	}

	println("err is not ErrSentinel")
}
