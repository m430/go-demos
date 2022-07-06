package main

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

// 函数f中的channel接收动作发生在main goroutine对channel发送动作完成之前
func main() {
	go f()
	c <- 5
	println(a)
}
