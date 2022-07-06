package unbuffed_chan

import "testing"

// 带缓冲区的channel当发送到缓冲区未满、接收操作在缓冲区非空的情况下都是异步的
// 也就是说对一个带缓冲区的channel，对其发送操作的goroutine不会阻塞，缓冲区满了才会阻塞
// 在缓存区未空的channel，对其进行接收操作不会阻塞，空了才会阻塞

// 消息队列
// 单收单发性能基准测试

var c1 chan string
var c2 chan string

func init() {
	c1 = make(chan string)
	go func() {
		for {
			<-c1
		}
	}()
	c2 = make(chan string)
	go func() {
		for {
			c2 <- "hello"
		}
	}()
}

func send(msg string) {
	c1 <- msg
}

func recv() {
	<-c2
}

func BenchmarkUnbuffedChan1To1Send(b *testing.B) {
	for i := 0; i < b.N; i++ {
		send("hello")
	}
}

func BenchmarkUnbuffedChan1To1Recv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recv()
	}
}
