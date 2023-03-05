package concurrence

import (
	"fmt"
	"time"
)

func Test_channel_rate_limit() {
	requests := make(chan int, 5)

	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond) // send a value every 200 millisecond

	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}

	}()

	burstRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)
	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request ", req, time.Now())
	}

}
