package goroutine

import (
	"fmt"
	"time"
)

func task(from string) {
	for i := 0; i < 20; i++ {
		fmt.Println("from ", from, ":", i)
	}
}

func Test() {
	task("direct")

	go task("from routine1")

	go func(msg string) {
		fmt.Println("inline task:", msg)
	}("inline")

	time.Sleep(time.Second)
	fmt.Println("done")
}
