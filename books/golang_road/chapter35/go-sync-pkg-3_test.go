package chapter35

import (
	"sync"
	"testing"
)

var cs1 = 0
var mu1 sync.Mutex
var cs2 = 0
var mu2 sync.RWMutex

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})
}

func BenchmarkWriteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu2.Lock()
			cs2++
			mu2.Unlock()
		}
	})
}

// 通过测试对比
// 并发量较小的情况下，互斥锁性能更好；随着并发量增大，互斥锁的竞争激烈导致加锁和解锁性能下降
// 读写锁的读锁性能并未随着并发量增大而变化，而写锁随着并发增大，性能降低

// 读写锁适合应用在具有一定并发量且读多写少的场合。大量并发读的情况下，多个goroutine可以同时持有读锁，从而减少锁竞争等待时间。
// 而互斥锁在同一时刻也只能有一个goroutine持有锁，其他goroutine会堵塞等待
