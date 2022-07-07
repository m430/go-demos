package main

import (
	"log"
	"sync"
	"time"
)

type Foo struct {
}

var once sync.Once
var instance *Foo

func GetInstance(id int) *Foo {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("goroutine-%d: caught a panic: %s", id, e)
		}
	}()
	log.Printf("goroutine-%d: enter GetInstance\n", id)
	once.Do(func() {
		instance = &Foo{}
		time.Sleep(3 * time.Second)
		log.Printf("goroutine-%d: the addr of instance is %p\n", id, instance)
		panic("panic in once.Do function")
	})
	return instance
}

// once.Do会等待f执行完毕后才返回，这期间其他执行once.Do函数的goroutine将会阻塞等待
// 第一个执行once.Do函数返回后，后续的goroutine再执行Do函数将不再执行f并立即返回
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			inst := GetInstance(i)
			log.Printf("goroutine-%d: the addr of instance  returned is %p\n", i, inst)
			wg.Done()
		}(i + 1)
	}
	time.Sleep(5 * time.Second)
	inst := GetInstance(0)
	log.Printf("goroutine-0: the addr of instance returned is %p\n", inst)
	wg.Wait()
	log.Printf("all goroutine exit\n")
}
