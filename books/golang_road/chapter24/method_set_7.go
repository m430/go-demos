package main

type T9 struct {
}

func (t T9) M1() {

}

func (t *T9) M2() {

}

type Interface9 interface {
	M1()
	M2()
}

type T10 T9
type Interface10 Interface9

//func main() {
//	var t T9
//	var pt *T9
//	var t1 T10
//	var pt1 *T10
//
//	DumpMethodSet(&t)
//	DumpMethodSet(&t1)
//
//	DumpMethodSet(&pt)
//	DumpMethodSet(&pt1)
//
//	DumpMethodSet((*Interface9)(nil))
//	DumpMethodSet((*Interface10)(nil))
//}

// 总结：基于接口类型创建的defined类型与原接口类型的方法集合是一致的，而基于自定义非接口类型创建的defined类型则并没有继承原类型的方法集合
// 所以即便原类型实现了接口，基于其创建的自定义类型也没有继承接口的实现。仍然需要重新实现所有方法
