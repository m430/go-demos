package main

import "time"

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * time.Duration(interval))
}

func spawn(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	go func() {
		f(args...)
		c <- struct{}{}
	}()
	return c
}

func main() {
	done := spawn(worker, 5)
	println("spawn a worker goroutine")
	<-done
	println("worker done")
}

// 总结：在内部创建一个goroutine并返回一个channel类型变量的函数，是Go中最常见的goroutine创建模式
// goroutine执行函数返回意味着goroutine退出，但是一些常驻后台服务程序可能会对goroutine有着优雅退出的要求
