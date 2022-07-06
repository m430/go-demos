package main

// 类型别名的方法集合
// 类型别名与原类型几乎是等价的，方法集合也完全相同

type T11 struct {
}

func (t T11) M1() {

}

func (t *T11) M2() {

}

type Interface11 interface {
	M1()
	M2()
}

type T12 = T11
type Interface12 = Interface11

func main() {
	var t T11
	var pt *T11
	var t1 T12
	var pt1 *T12

	DumpMethodSet(&t)
	DumpMethodSet(&t1)

	DumpMethodSet(&pt)
	DumpMethodSet(&pt1)

	DumpMethodSet((*Interface11)(nil))
	DumpMethodSet((*Interface12)(nil))
}
