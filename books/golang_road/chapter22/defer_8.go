package main

import "fmt"

func fooo1() {
	s1 := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a)
	}(s1)

	s1 = []int{3, 2, 1}
	_ = s1
}

func fooo2() {
	s1 := []int{1, 2, 3}
	defer func(p *[]int) {
		fmt.Println(*p)
	}(&s1)

	s1 = []int{3, 2, 1}
	_ = s1
}

func main() {
	fooo1()
	fooo2()
}

// 总结：defer函数在外部函数调用后执行，如果defer函数调用传递的是外部作用域指针，则需要将外部作用域执行完毕的指针值作为defer函数的参数
// 如果defer函数传递的不是指针，则以压入defer函数栈时刻的值为参数进行执行
