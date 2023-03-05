package channel

import (
	"fmt"
)

// to-do: interesting, how to fix this error??
func Test_channel_range_over_without_close() {
	c := make(chan string, 2)
	//done := make(chan bool)
	c <- "one"
	c <- "two"
	close(c) // there is an error when we comment this line

	/*
		for i := 0; i < 2; i++ {
			fmt.Println("received:", <-c)
		}*/

	for e := range c {
		fmt.Println("received:", e)
	}

	//<-done

}
