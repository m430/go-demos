package main

import (
	"log"
	"sync"
	"time"
)

type foo struct {
	n int
	sync.Mutex
}

// 创建了两个goroutine：g2和g3，g3阻塞在加锁操作上
func main() {
	f := foo{n: 17}

	go func(f foo) {
		for {
			log.Println("g2: try to lock foo...")
			f.Lock()
			log.Println("g2: lock foo ok")
			time.Sleep(3 * time.Second)
			f.Unlock()
			log.Println("g2: unlock foo ok")
		}
	}(f)

	f.Lock()
	log.Println("g1: lock foo ok")

	// 在Mutex首次使用后复制其值
	go func(f foo) {
		for {
			log.Println("g3: try to lock foo...")
			f.Lock()
			log.Println("g3: lock foo ok")
			time.Sleep(5 * time.Second)
			f.Unlock()
			log.Println("g3: unlock foo ok")
		}
	}(f)

	time.Sleep(1000 * time.Second)
	f.Unlock()
	log.Println("g1: unlock foo ok")
}
