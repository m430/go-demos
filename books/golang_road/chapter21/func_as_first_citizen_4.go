package main

import "fmt"

// 函数式编程
func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	timesTwo := partialTimes(2)
	timesThree := partialTimes(3)
	timesFour := partialTimes(4)
	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
	fmt.Println(timesFour(5))
}

// 总结：上面例子实现了函数柯里化，使用了闭包的特性。闭包是在函数内部定义的匿名函数，并允许该匿名函数访问定义它的外部函数作用域
