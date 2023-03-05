package concurrence

import (
	"fmt"
	"time"
)

func Test_concurrence_timer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // this line is blocked until time event hits
	fmt.Println("timers is fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C // this line is blocked until time event comes. however
		//  the time event will never come because we stop the timer at line 19. The line 19 results in that this goroutine is stuck here for ever
		fmt.Println("timer2 is fired")
	}()
	stop2 := timer2.Stop() // time can be stopped before it is fired
	if stop2 {
		fmt.Println("timer2 is stopped")
	}

	time.Sleep(3 * time.Second)

}
