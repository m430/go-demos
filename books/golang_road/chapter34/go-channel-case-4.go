package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct {
}

func worker(i int, quit <-chan signal) {
	fmt.Printf("worker-%d: is working...\n", i)
LOOP:
	for {
		select {
		default:
			// 模拟worker工作
			time.Sleep(1 * time.Second)
		case <-quit:
			break LOOP
		}
	}
	fmt.Printf("worker-%d: works done\n", i)

}

func spawnGroup(f func(int, <-chan signal), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker-%d: start to work...\n", i)
			f(i, groupSignal)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

// 一对多通知信号，通知一组worker goroutine退出
func main() {
	println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	fmt.Println("the group of workers start to work...")
	time.Sleep(5 * time.Second)
	// 通知workers退出
	fmt.Println("notify the group of workers to exit...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}

// 总结：关闭一个无缓冲的channel会让所有阻塞在该channel上的接收操作返回。从而实现一对多的广播机制。
