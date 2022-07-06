package main

import (
	"fmt"
	"time"
)

type counter struct {
	c chan int
	i int
}

var cter counter

func InitCounter() {
	cter = counter{
		c: make(chan int),
	}

	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	fmt.Println("counter init ok")
}

func Increase() int {
	return <-cter.c
}

func init() {
	InitCounter()
}

// 用于替代锁机制
// 将Counter操作全部交给一个独立的goroutine处理，并通过无缓冲channel的同步阻塞特性实现计数器的控制
// 这样其他goroutine通过increase函数试图增加计数器的动作实质上转化为一次无缓冲channel的接收动作
// 这种并发设计符合Go语言倡导的"不要通过共享内存来通信，而应该通过通信来共享内存"
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
		}(i)
	}
	time.Sleep(5 * time.Second)
}
