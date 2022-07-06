package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ": idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ": bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\tgoroutine-", id, ": xRayCheck ok\n")
	return xRayCheckTmCost
}

func newAirportSecurityCheck(id string, queue <-chan struct{}) {
	go func(id string) {
		println("goroutine-", id, ": airportSecurityCheckChannel is ready...")

		queue3, quit3, result3 := start(id, xRayCheck, nil)
		queue2, quit2, result2 := start(id, bodyCheck, queue3)
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					print("goroutine-", id, ": airportSecurityCheckChannel time cost:", total, "\n")
					print("goroutine-", id, ": airportSecurityCheckChannel closed\n")
					return
				}
				queue1 <- v
			}
		}
	}(id)
}

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)
	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()
	return queue, quit, result
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	passengers := 30
	queue := make(chan struct{}, 30)
	newAirportSecurityCheck("channel1", queue)
	newAirportSecurityCheck("channel2", queue)
	newAirportSecurityCheck("channel3", queue)

	time.Sleep(5 * time.Second)
	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(5 * time.Second)
	close(queue)
	time.Sleep(1000 * time.Second)
}
