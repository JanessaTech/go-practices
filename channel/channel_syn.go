package channel

import (
	"fmt"
	"time"
)

func task(done chan bool) {
	fmt.Println("starting ....")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
	done <- true // send a notification at the end of go routine
}

func Test_channel_syn() {
	done := make(chan bool, 1) // create a channel

	go task(done) //execute go routine

	<-done // notice that code will wait here util there is something in done channel
}
