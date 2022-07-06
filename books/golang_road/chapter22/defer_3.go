package main

import "fmt"

//#include <stdio.h>
//void crash() {
//	int *q = NULL;
//  (*q) = 15000;
//  printf("%d\n", *q)
//}

import "C"

func bar1() {
	C.crash()
}

func foo1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered from a panic: ", e)
		}
	}()
	bar1()
}

func main() {
	foo1()
	fmt.Println("main exit normally")
}
