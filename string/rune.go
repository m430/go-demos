package main

import "fmt"

func main() {
	var b = "Go语言"
	fmt.Println(len(b))

	for i, v := range b {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}
}
