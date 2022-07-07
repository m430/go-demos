package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct {
}

var ready bool

func worker(i int) {
	fmt.Printf("worker-%d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker-%d: works done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				groupSignal.L.Lock()
				for !ready {
					groupSignal.Wait()
				}
				groupSignal.L.Unlock()
				fmt.Printf("worker-%d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

// sync.Cond实例初始化需要一个满足sync.Locker接口的实例，通常用互斥锁sync.Mutex
// sync.Cond需要这个互斥锁来保护用作条件的数据，各个等待条件成立的goroutine在加锁后判断条件是否成立
// 如果不成立，则调用sync.Cond的Wait方法进入等待状态。
// main goroutine将ready设为true并调用sync.Cond的Broadcast方法后，各个阻塞的goroutine将被唤醒继续执行后续逻辑
func main() {
	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroup(worker, 5, groupSignal)

	time.Sleep(5 * time.Second) // 模拟ready前准备工作
	fmt.Println("the group of workers start to work")

	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
