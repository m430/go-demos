package main

import (
	"fmt"
	"time"
)

type signal struct {
}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

// 一对一通知信号
func main() {
	println("start a worker...")
	c := spawn(worker)
	// main goroutine阻塞在c的接收上，直到spawn里的goroutine发送给c一个signal
	<-c
	fmt.Println("worker work done!")
}
