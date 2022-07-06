package buffed_chan

import "testing"

var c1 chan string
var c2 chan string

func init() {
	c1 = make(chan string, 10)
	go func() {
		for {
			<-c1
		}
	}()
	c2 = make(chan string, 10)
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

func BenchmarkBuffedChan1To1Send(b *testing.B) {
	for i := 0; i < b.N; i++ {
		send("hello")
	}
}

func BenchmarkBuffedChan1To1Recv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recv()
	}
}
