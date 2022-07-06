package main

type Interface2 interface {
	M1()
	M2()
}

type T2 struct {
}

func (t T2) M1() {

}

func (t *T2) M2() {

}

//func main() {
//	var t T2
//	var pt *T2
//	DumpMethodSet(&t)
//	DumpMethodSet(&pt)
//	DumpMethodSet((*Interface2)(nil))
//}
