package main

type Interface3 interface {
	M1()
	M2()
}

type T3 struct {
	Interface3
}

func (t T3) M3() {

}

//func main() {
//	DumpMethodSet((*Interface3)(nil))
//	var t T3
//	var pt *T3
//	DumpMethodSet(&t)
//	DumpMethodSet(&pt)
//}
