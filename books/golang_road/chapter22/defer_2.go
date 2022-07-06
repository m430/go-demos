package main

import "fmt"

func bar() {
	fmt.Println("raise a panic")
	panic(-1)
}

func foo() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered from a panic")
		}
	}()
	bar()
}

func main() {
	foo()
	fmt.Println("main exit normally")
}

// 总结：defer函数在出现panic情况下依然会被调用，所以竟然在defer函数中处理panic
