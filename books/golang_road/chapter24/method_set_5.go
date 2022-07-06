package main

type Interface5 interface {
	M1()
	M2()
}

type T5 struct {
	Interface5
}

func (t T5) M1() {
	println("T5's M1")
}

type S struct {
}

func (s S) M1() {
	println("S's M1")
}

func (s S) M2() {
	println("S's M2")
}

//func main() {
//	var t = T5{
//		Interface5: S{},
//	}
//
//	t.M1()
//	t.M2()
//}

// 如果结构体中嵌入多个接口的方法集合存在交集时，Go选择方法次序如下：
// 1.优选选择结构体自身实现方法
// 2. 如果结构体未实现该方法，查找结构体中的嵌入接口类型的方法集合中是否有该方法，如果有则提升为结构体的方法
// 3. 如果嵌入多个接口类型且这些接口类型方法集合存在交集，那么Go编译器会报错
