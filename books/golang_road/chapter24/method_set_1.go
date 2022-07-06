package main

type Interface interface {
	M1()
	M2()
}

type T struct {
}

func (t T) M1() {

}

func (t *T) M2() {

}

//func main() {
//	var t T
//	var pt *T
//	var i Interface
//	i = t
//	i = pt
//}

// 总结：方法集合决定接口实现
