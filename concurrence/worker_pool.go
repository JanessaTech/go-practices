package concurrence

import (
	"fmt"
	"time"
)

const numJobs = 5

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", workerId, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", workerId, "finished job ", j)
		results <- j * 2
	}
}

func Test_concurrence_worker_pool() {

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	/*//no error in this solution
	for a := 1; a <= numJobs; a++ {
		<-results
	}*/

	// code has some error because the channel results is not closed
	for r := range results {
		fmt.Println("result:", r)
	}
}
