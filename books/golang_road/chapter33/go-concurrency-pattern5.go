package main

import (
	"time"
)

func worker(j int) {
	time.Sleep(time.Second * time.Duration(j))
}

func spawn(f func(int)) chan string {
	quit := make(chan string)
	go func() {
		var job chan int
		for {
			select {
			case j := <-job:
				f(j)
			case <-quit:
				quit <- "ok"
			}
		}
	}()
	return quit
}

// notify-and-wait模式
// 前面几个场景，goroutine创建者都是被动等待新的goroutine退出，很多时候，创建者需要主动通知那些新goroutine退出
// 1. 通知并等待一个goroutine退出
func main() {
	quit := spawn(worker)
	println("spawn a worker goroutine")

	time.Sleep(5 * time.Second)

	println("notify the worker to exit...")
	quit <- "exit"

	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case status := <-quit:
		println("worker done:", status)
	case <-timer.C:
		println("wait worker exit timeout")
	}
}
