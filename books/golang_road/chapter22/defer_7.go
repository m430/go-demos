package main

import "fmt"

func fooo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func fooo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func fooo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	fmt.Println("fooo1 result:")
	fooo1()
	fmt.Println("fooo2 result:")
	fooo2()
	fmt.Println("fooo3 result:")
	fooo3()
}

// 总结：defer函数栈将以先进后出的方式执行。
