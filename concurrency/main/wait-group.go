package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d ended\n", id)
}

func main() {

	var waitGroup sync.WaitGroup

	for i := 1; i < 6; i++ {
		waitGroup.Add(i)
		i := i
		go func() {
			defer waitGroup.Done()
			worker(i)
		}()
	}

	waitGroup.Wait()
}
