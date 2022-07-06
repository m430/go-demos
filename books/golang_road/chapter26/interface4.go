package main

import "fmt"

type T struct {
	n int
	s string
}

func (t T) M1() {

}

func (t T) M2() {

}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T{
		n: 17,
		s: "hello, interface",
	}
	var ei interface{}
	ei = t

	var i NonEmptyInterface
	i = t
	fmt.Println(ei)
	fmt.Println(i)
}
