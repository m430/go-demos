package main

import "fmt"

func printNilInterface() {
	var i interface{}
	var err error
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
	println(" ")
}

func printEmptyInterface() {
	var eif1 interface{}
	var eif2 interface{}
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)

	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)

	println(" ")
}

type T int

func (t T) Error() string {
	return "bad error"
}

func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	println("")
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
}

//func main() {
//	printNilInterface()
//	println(" ")
//
//	printEmptyInterface()
//	println(" ")
//
//	printNonEmptyInterface()
//	println(" ")
//
//	printEmptyInterfaceAndNonEmptyInterface()
//	println(" ")
//}
