package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	id int
}

func worker(workerID int, ch chan job, wg *sync.WaitGroup) {

	defer wg.Done()

	for jb := range ch {

		fmt.Println("Worker", workerID, "is executing job", jb.id)
		
		time.Sleep(1*time.Second)
	}
}

func main() {

	var wg sync.WaitGroup

	ch := make(chan job, 100)

	var n int
	var j int

	fmt.Println("Enter the job count:")
	fmt.Scan(&n)

	fmt.Println("Enter the worker count:")
	fmt.Scan(&j)

	// workers
	for i := 1; i <= j; i++ {

		wg.Add(1)

		go worker(i, ch, &wg)
	}

	// jobs
	for i := 1; i <= n; i++ {

		ch <- job{id: i}
	}

	close(ch)

	wg.Wait()
}
