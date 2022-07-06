package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(j int) {
	time.Sleep(time.Second * time.Duration(j))
}

func spawnGroup(n int, f func(int)) chan struct{} {
	quit := make(chan struct{})
	job := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			// 保证wg在goroutine退出前执行
			defer wg.Done()
			name := fmt.Sprintf("worker-%d", i)
			for {
				j, ok := <-job
				if !ok {
					println(name, "done")
					return
				}
				worker(j)
			}
		}(i)
	}

	go func() {
		<-quit
		// 广播给所有新的goroutine
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()
	return quit
}

// notify-and-wait模式
// 前面几个场景，goroutine创建者都是被动等待新的goroutine退出，很多时候，创建者需要主动通知那些新goroutine退出
// 1. 通知并等待多个goroutine退出
func main() {
	quit := spawnGroup(5, worker)
	println("spawn a group of workers")

	time.Sleep(5 * time.Second)

	println("notify the worker group to exit...")
	quit <- struct{}{}

	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case <-quit:
		println("worker group done")
	case <-timer.C:
		println("wait worker group exit timeout")
	}
}
