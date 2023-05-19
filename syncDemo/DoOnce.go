package syncdemo

import (
	"fmt"
	"sync"
)

var onlyOnce sync.Once

func doConfig() {
	onlyOnce.Do(func() {
		fmt.Printf("Read configuration file from I/O disk")
	})
	fmt.Println("done")

}

func DoOnceDemo() {
	go doConfig()
	go doConfig()
	go doConfig()

	// keep  running until ctrl-c
	for {
	}
}
