package channel

import (
	"fmt"
)

func Test_channel_close() {
	c := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for { // endless loop in for
			j, more := <-c //more is true if c is alive
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("channel is closed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		c <- i
		fmt.Println("send: ", i)
	}
	close(c)

	<-done // notice that code will wait here util there is something in done channel. It is like the usage of latch in java

}
