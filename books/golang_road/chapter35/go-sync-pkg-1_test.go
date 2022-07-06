package chapter35

import (
	"sync"
	"testing"
)

var cs = 0
var mu sync.Mutex
var c = make(chan struct{}, 1)

func criticalSectionSyncByMutex() {
	mu.Lock()
	cs++
	mu.Unlock()
}

func criticalSectionSyncByChan() {
	c <- struct{}{}
	cs++
	<-c
}

// sync.Mutex性能比channel实现高出两倍
func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByChan()
	}
}

// 虽然Go语言倡导"不要通过共享内存内通信，而应该通过通信来共享内存"，但是在一些场景中还是用sync包来实现比较适合
// 1. 需要高性能的临界区同步机制场景
// 2. 不想转移结构体对象所有权，但又想保证结构体的内部状态数据的同步访问的场景
