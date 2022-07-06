package main

import "fmt"

func dump(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

//func main() {
//	s := []interface{}{"Tony", "John", "Jim"}
//	dump(s...)
//}
