package main

import (
	"bytes"
	"sync"
	"testing"
)

// sync.Pool是一个数据对象缓存池
// sync.Pool是并发安全的，可以被多个goroutine同时使用
// 放入缓存池中的数据对象的生命是暂时的，随时可能被垃圾回收掉
// 缓存池中的对象是可以重复使用的，这样可以一定程度降低数据对象重新分配的频度，减轻GC压力
// sync.Pool为每个P（goroutine调度模型中的P）单独建立一个local缓存池，进一步降低高并发下对锁的争抢

// sync.Pool一个典型的应用就是建立像bytes.Buffer这样类型的临时缓存对象池
// 但是，sync.Pool的Get方法从缓存池中挑选bytes.Buffer数据对象时并未考虑该数据对象是否满足调用者需求
// 因此一旦返回对象是刚刚被撑大的，并且即将长期用于处理一些小数据时，这个Buffer对象所占用的大内存将长时间得不到释放。
// 目前Go标准库采用两种方式来缓解这个问题：
// 1. 限制要放回缓存池中的数据对象大小
// 2. 建立多级缓存池（例如http包处理http2数据时，预先创建多个不同大小的缓存池）
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func writeBufFromPool(data string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(data)
	bufPool.Put(b)
}

func writeBufFromNew(data string) *bytes.Buffer {
	b := new(bytes.Buffer)
	b.WriteString(data)
	return b
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBufFromNew("hello")
	}
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		writeBufFromPool("hello")
	}
}
