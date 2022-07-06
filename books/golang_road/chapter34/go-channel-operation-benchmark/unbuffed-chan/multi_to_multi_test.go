package unbuffed_chan

import "testing"

var c1 chan string
var c2 chan string

func init() {
	c1 = make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				<-c1
			}
		}()
		go func() {
			for {
				c1 <- "hello"
			}
		}()
	}

	c2 = make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				c2 <- "hello"
			}
		}()
		go func() {
			<-c2
		}()
	}
}

func send(msg string) {
	c1 <- msg
}

func recv() {
	<-c2
}

func BenchmarkUnbuffedChanNToNSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		send("hello")
	}
}

func BenchmarkUnbuffedChanNToNRecv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recv()
	}
}
