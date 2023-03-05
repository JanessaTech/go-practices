package channel

import "fmt"

func Test_channel_range_over() {
	c := make(chan string, 2)
	c <- "one"
	c <- "two"
	close(c)

	for e := range c {
		fmt.Println("received:", e)
	}
}
