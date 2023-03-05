package channel

import "fmt"

func Test_channel_basic() {
	messsage := make(chan string)

	go func(msg string) { // define a inline function, then call it
		messsage <- msg
	}("hello")

	received := <-messsage
	fmt.Println(received)
}
