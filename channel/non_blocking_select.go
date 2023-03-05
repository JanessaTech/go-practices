package channel

import "fmt"

func Test_channel_select_nonblocking() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default: // hit here if no messsage in the channel messages
		fmt.Println("no message received")
	}

	msg := "Hello world"
	select {
	case messages <- msg:
		fmt.Println("send message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}

}
