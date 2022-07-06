package main

import "fmt"

// 延续传递式
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Printf("%d\n", factorial(5))
}

// 总结：上面例子实现了函数柯里化，使用了闭包的特性。闭包是在函数内部定义的匿名函数，并允许该匿名函数访问定义它的外部函数作用域
