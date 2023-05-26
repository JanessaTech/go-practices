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

func subTask1Routine(ctx context.Context) {
	fmt.Println("enter subTask1Routine...")
	for { // for loop running endless until it hits return
		select {
		case <-ctx.Done(): // cancel() in withCancel is called or deadline is reached
			err := ctx.Err()
			fmt.Println("context in subTask1Routine is stopped. reason:", err)
			return
		default:
			fmt.Println("do something else for sub task1...")
			time.Sleep(1 * time.Second)
		}
	}
}
func subTask2Routine(ctx context.Context) {
	fmt.Println("enter subTask2Routine...")
	for { // for loop running endless until it hits return
		select {
		case <-ctx.Done(): // cancel() in withCancel is called or deadline is reached
			err := ctx.Err()
			fmt.Println("context in subTask2Routine is stopped. reason:", err)
			return
		default:
			fmt.Println("do something else for sub task2...")
			time.Sleep(1 * time.Second)
		}
	}
}

func withCancel() {
	ctx := context.Background()
	fmt.Println("starting main task...")
	newCtx, cancel := context.WithCancel(ctx)

	defer cancel() //it isn’t necessarily required

	go subTask1Routine(newCtx)
	go subTask2Routine(newCtx)

	time.Sleep(3 * time.Second)
	cancel() // call cancel() to signal all goroutines that newCtx is closed

	time.Sleep(1 * time.Second)
	fmt.Println("main task is finished")
}

func withDeadline() {
	ctx := context.Background()
	deadline := time.Now().Add(2 * time.Second)
	fmt.Println("starting main task...")
	newCtx, cancel := context.WithDeadline(ctx, deadline)

	defer cancel() //it isn’t necessarily required

	go subTask1Routine(newCtx)
	go subTask2Routine(newCtx)

	time.Sleep(3 * time.Second)
	cancel() // call cancel() to signal all goroutines that newCtx is closed

	time.Sleep(1 * time.Second)
	fmt.Println("main task is finished")
}

func withTimeout() {
	ctx := context.Background()
	fmt.Println("starting main task...")
	newCtx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel() //it isn’t necessarily required

	go subTask1Routine(newCtx)
	go subTask2Routine(newCtx)

	time.Sleep(3 * time.Second)
	cancel() // call cancel() to signal all goroutines that newCtx is closed

	time.Sleep(1 * time.Second)
	fmt.Println("main task is finished")

}

// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
// https://zhuanlan.zhihu.com/p/611879579
func Main() {
	//storeValueInContext()
	withCancel()
	//withDeadline()
	//withTimeout()
}
