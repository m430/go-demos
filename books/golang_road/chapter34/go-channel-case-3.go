package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct {
}

func worker(i int) {
	fmt.Printf("worker-%d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker-%d: works done\n", i)

}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker-%d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

// 一对多通知信号
// main goroutine里面创建5个worker goroutine,启动后会阻塞在groupSignal上。
// main goroutine通过close(groupSignal)向所有worker goroutine广播开始工作
// 所有worker goroutine收到groupSignal信号后开始开始工作，跑完后。
// spawnGroup还启动了一个监听goroutine，当waitGroup跑完后，会向c的channel中发送一个signal
// main goroutine中的c channel接收到信号后，继续往下执行
func main() {
	println("start a group of workers...")
	groupSinal := make(chan signal)
	c := spawnGroup(worker, 5, groupSinal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSinal)
	<-c
	fmt.Println("the group of workers work done!")
}

// 总结：关闭一个无缓冲的channel会让所有阻塞在该channel上的接收操作返回。从而实现一对多的广播机制。
