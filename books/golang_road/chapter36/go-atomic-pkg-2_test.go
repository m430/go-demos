package chapter36

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	sync.RWMutex
	data string
}

func BenchmarkRWMutexSet(b *testing.B) {
	config := Config{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Lock()
			config.data = "hello"
			config.Unlock()
		}
	})
}

func BenchmarkRWMutexGet(b *testing.B) {
	config := Config{data: "hello"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.RLock()
			_ = config.data
			config.RUnlock()
		}
	})
}

func BenchmarkAtomicSet(b *testing.B) {
	var config atomic.Value
	c := Config{data: "hello"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{data: "hello"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}

// 根据上述测试，得出结论
// 1. 利用原子操作的无锁并发写的性能随着并发量增大而小幅下降
// 2. 利用原子操作的无锁并发读性能随着并发量增大有持续提升趋势，并且性能约为读锁的100倍

// 所以，atomic包更适合一些对性能十分敏感、并发量较大且读多写少的场合
// 但是，atomic原子操作可用来同步的范围有较大限制，仅是一个整形变量或自定义类型，如果是复杂临界区数据进行同步，依然首选sync包
