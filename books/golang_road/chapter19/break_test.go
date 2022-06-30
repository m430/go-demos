package chapter19

import (
	"fmt"
	"testing"
	"time"
)

// break
func TestBreak1(t *testing.T) {
	exit := make(chan interface{})

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				// 实际上跳出的是select而不是for循环
				break
			}
		}
		fmt.Println("exit!")
	}()

	time.Sleep(3 * time.Second)
	exit <- struct{}{}

	// 等待子goroutine退出
	time.Sleep(3 * time.Second)
}

// break label
func TestBreak2(t *testing.T) {
	exit := make(chan interface{})

	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				// 这里break会跳出到loop内的for循环
				break loop
			}
		}
		fmt.Println("exit!")
	}()

	time.Sleep(3 * time.Second)
	exit <- struct{}{}

	// 等待子goroutine退出
	time.Sleep(3 * time.Second)
}

// 总结1：break语句和C语言中不同，Go语言规定break语句结束执行并跳出的是同一函数内break语句所在的最内层的for、switch、select语句。
// 总结2：带label的continue和break可以让深层循环终止外层循环或跳转到外层循环继续执行的能力
