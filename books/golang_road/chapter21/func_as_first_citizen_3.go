package main

import "fmt"

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func main() {
	var i BinaryAdder = MyAdderFunc(MyAdd)
	fmt.Println(i.Add(5, 6))
}

// 总结：函数也可以进行类型转换，MyAdd直接赋值给BinaryAdder接口是不行的，但是可以用已经实现BinaryAdder接口的MyAdderFunc进行类型转换
