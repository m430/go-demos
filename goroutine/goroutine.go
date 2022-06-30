package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("电脑核心数是%d\n", runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	for i := 1; i < 100; i++ {
	//		time.Sleep(1 * time.Millisecond)
	//		fmt.Println("A:", i)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 100; i++ {
	//		time.Sleep(1 * time.Millisecond)
	//		fmt.Println("B:", i)
	//	}
	//}()
	//
	//wg.Wait()
}
