package concurrence

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Test_concurrence_atomic_counter() {
	var cnt uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&cnt, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("cnt:", cnt)
}
