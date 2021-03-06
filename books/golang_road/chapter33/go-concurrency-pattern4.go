package main

import (
	"fmt"
	"sync"
	"time"
)

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

func spawnGroup(n int, f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			name := fmt.Sprintf("worker-%d:", i)
			f(args...)
			println(name, "done")
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		c <- struct{}{}
	}()
	return c
}

// 支持超时机制的等待
func main() {
	done := spawnGroup(5, worker, 30)
	println("spawn a group of workers")

	// 通过一个定时器设置了超时等待时间，并通过select原语同时监听timer.C和done这两个channel，那个先返回就执行那个case分支
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	select {
	case <-timer.C:
		println("wait group workers exit timeout!")
	case <-done:
		println("group workers done")
	}
}
