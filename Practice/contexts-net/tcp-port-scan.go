package main 

//NOTE: 
/*
->The OS is not able to spin out 65350 tcp connections at the same time, 
->A Worker Pool solves this by opening only 500 connections at a time.
->Contexts are used here to prevent the go routinees from being hanging while the connections are refused and fallback.
*/

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(
	workerID int,
	ports <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for port := range ports {

		ctx, cancel := context.WithTimeout(
			context.Background(),
			500*time.Millisecond,
		)

		address := fmt.Sprintf("127.0.0.1:%d", port)

		var d net.Dialer

		conn, err := d.DialContext(ctx, "tcp", address)

		cancel()

		if err != nil {
			continue
		}

		results <- port

		conn.Close()
	}
}

func main() {

	var wg sync.WaitGroup

	ports := make(chan int, 500)

	results := make(chan int)

	done := make(chan struct{})

	for workers := 1; workers <= 500; workers++ {
		wg.Add(1)

		go worker(
			workers,
			ports,
			results,
			&wg,
		)
	}

	go func() {
		for port := 1; port <= 65350; port++ {
			ports <- port
		}

		close(ports)
	}()

	go func() {
		for res := range results {
			fmt.Printf("Active Port: %d\n", res)
		}

		close(done)
	}()

	wg.Wait()

	close(results)

	<-done
}
