package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

// errors.As会沿着err2所在的错误链向上找到被包装到最深处的错误值，并将err2与其类型*MyError成功匹配
func main() {
	var err = &MyError{"my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)

	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}

	println("MyError is not on the chain of err2")
}
