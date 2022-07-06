package main

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func main() {
	var t T
	println(t.a)

	t.M1()
	println(t.a)

	t.M2()
	println(t.a)
}

// 总结1：method的本质其实就是把receiver当做第一个参数的函数
// 总结2：如果要对类实例进行修改，receiver类型选择指针类型
