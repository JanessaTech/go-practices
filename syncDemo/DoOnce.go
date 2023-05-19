package syncdemo

import (
	"fmt"
	"sync"
)

var onlyOnce sync.Once

func doOnceCall() {
	onlyOnce.Do(func() {
		fmt.Printf("It is called for the first time")
	})
	fmt.Println("done")

}

func DoOnceDemo() {
	go doOnceCall()
	go doOnceCall()
	go doOnceCall()

	// keep  running until ctrl-c
	for {
	}
}
