package contextusage

import (
	"context"
	"fmt"
	"time"
)

func storeValue(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "myName", "Janessa")
	return ctx
}
func getValue(ctx context.Context) {
	value := ctx.Value("myName")
	fmt.Println("myName =", value)
}

func storeValueInContext() {
	ctx := context.Background()
	ctx = storeValue(ctx)
	getValue(ctx)
}

func DeadlinesUsingWithTimeout() {
	fmt.Println("starting ...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("timed out")
				err := c.Err()
				fmt.Println(err)
				return
			default:
				fmt.Println("doing your tasks ...")
			}
			time.Sleep(500 * time.Millisecond)

		}
	}(ctx)

	// keep  running until ctrl-c
	for {
	}
}

// https://zhuanlan.zhihu.com/p/611879579
func Main() {
	storeValueInContext()
	//DeadlinesUsingWithTimeout()
}
