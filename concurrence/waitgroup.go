package concurrence

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Println("worker ", id, " is starting..")
	time.Sleep(time.Second)
	fmt.Println("worker ", id, " is finished")
}

func execute_defer(i int, wg *sync.WaitGroup) {
	fmt.Println("executing defer", i, " ...")
	wg.Done() // decrease by 1
}

func Test_concurrence_waitgroup() {
	var wg sync.WaitGroup // it looks like the latch in Java

	for i := 0; i < 5; i++ {
		wg.Add(1) // increase by 1
		i := i
		go func() {
			//defer wg.Done()
			defer execute_defer(i, &wg) // execute this line after work(i) is done
			work(i)
		}()
	}
	wg.Wait() // wait until wg is 0
}
