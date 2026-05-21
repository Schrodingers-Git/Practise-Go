package main

/*
	fan out the workers to calculate the factorial,
	fan in the results from multiple workers.
*/

import (
	"fmt"
	"sync"
	"time"
)

// fan out workers
func worker_fan_out(start int, end int, ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	var res int
	res = 1

	for i := start; i <= end; i++ {
		res *= i
	}

	ch <- res
}

// fan in worker
func worker_fan_in(res *int, mu *sync.Mutex, ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := range ch {

		mu.Lock()
		*res *= i
		mu.Unlock()

		time.Sleep(1 * time.Second)
	}
}

func main() {

	var wg sync.WaitGroup

	var n int

	fmt.Println("Enter the number to find the factorial : ")
	fmt.Scan(&n)

	ch := make(chan int)

	m := n / 5

	// fan out
	for i := 0; i < m; i++ {

		wg.Add(1)

		go worker_fan_out((i*5)+1, (i*5)+5, ch, &wg)
	}

	res := 1

	var mu sync.Mutex

	go func() {

		wg.Wait()

		close(ch)

	}()

	// fan in let number of workers - 5
	for i := 1; i <= 5; i++ {

		wg.Add(1)

		go worker_fan_in(&res, &mu, ch, &wg)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Result is : ", res)
}
