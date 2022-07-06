package main

import (
	"fmt"
	"time"
)

// 对没有初始化的channel（nil channel）进行读写操作将会发生阻塞
// 1. 前5s，select一直处于阻塞状态
// 2. 第5s，c1返回一个5后被关闭，select语句的第一个case分支被执行，然后继续for循环
// 3. c1被关闭，由于一个被关闭的channel接收输入不会被阻塞，所以新一轮的case分支获取c1的零值，会一直输出0
// 4. 2s后，c2被写入一个7，case的第二个分支被执行，然后退出for循环
func main() {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		c2 <- 7
		close(c2)
	}()

	var ok1, ok2 bool
	for {
		select {
		case x := <-c1:
			ok1 = true
			fmt.Println(x)
		case x := <-c2:
			ok2 = true
			fmt.Println(x)
		}

		if ok1 && ok2 {
			break
		}
	}

	fmt.Println("program end")
}
