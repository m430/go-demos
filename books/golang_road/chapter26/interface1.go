package main

import (
	"errors"
)

type MyError struct {
	error
}

var ErrBad = MyError{errors.New("bad error")}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

//func main() {
//	e := returnsError()
//	// nil error != nil
//	if e != nil {
//		fmt.Printf("error: %+v\n", e)
//		return
//	}
//	fmt.Println("ok")
//}
