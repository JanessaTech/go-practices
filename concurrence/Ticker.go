package concurrence

import (
	"fmt"
	"time"
)

func Test_concurrence_ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("go routine is finished")
				return
			case t := <-ticker.C:
				fmt.Println("Ticker at: ", t)
			}
		}

	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
