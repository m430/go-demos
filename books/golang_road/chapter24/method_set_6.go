package main

// 结构体中嵌入结构体
type T6 struct {
}

func (t T6) T6M1() {
	println("T6' M1")
}

func (t T6) T6M2() {
	println("T6's M2")
}

func (t *T6) PT6M3() {
	println("PT6's M3")
}

type T7 struct {
}

func (t T7) T7M1() {
	println("T7' M1")
}

func (t T7) T7M2() {
	println("T7's M2")
}

func (t *T7) PT7M3() {
	println("PT7's M3")
}

type T8 struct {
	T6
	*T7
}

//func main() {
//	t := T8{
//		T6: T6{},
//		T7: &T7{},
//	}
//
//	println("call method through t:")
//	t.T6M1()
//	t.T6M2()
//	t.PT6M3()
//	t.T7M1()
//	t.T7M2()
//	t.PT7M3()
//
//	println("\ncall method through pt:")
//	pt := &t
//	pt.T6M1()
//	pt.T6M2()
//	pt.PT6M3()
//	pt.T7M1()
//	pt.T7M2()
//	pt.PT7M3()
//
//	var t6 T6
//	var pt6 *T6
//	DumpMethodSet(&t6)
//	DumpMethodSet(&pt6)
//
//	var t7 T7
//	var pt7 *T7
//	DumpMethodSet(&t7)
//	DumpMethodSet(&pt7)
//
//	DumpMethodSet(&t)
//	DumpMethodSet(&pt)
//}
