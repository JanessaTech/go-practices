package channel

import "fmt"

func Test_channel_buffer() {
	messsage := make(chan string, 2)

	messsage <- "hello"
	messsage <- "world"

	fmt.Println(<-messsage)
	fmt.Println(<-messsage)
}
