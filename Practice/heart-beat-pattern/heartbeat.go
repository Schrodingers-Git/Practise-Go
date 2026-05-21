package main

import (
	"fmt"
	"time"
)

func heartbeat(interval time.Duration, stop chan struct{}) <-chan struct{} {

	//hearbeat channel for the consumers of the heartbeat to know the ticks
	heartbeat := make(chan struct{})

	go func() {

		//create ticker
		//->it internally creates a ticker type/struct which has a channel(.C is used to access that channel)
		// that keeps sending time values for every interval
		//->stop is separate channel,used to monitor on when to stop the ticker
		ticker := time.NewTicker(interval)

		//IMPORTANT:
		// ticker internally allocates runtime timer resources
		// so always stop it properly
		defer ticker.Stop()

		for {

			select {

			case <-ticker.C:
				heartbeat <- struct{}{}

			case <-stop:
				close(heartbeat)
				return
			}
		}
	}()

	return heartbeat
}

func main() {

	stop := make(chan struct{})

	heartbeatChannel := heartbeat(500*time.Millisecond, stop)

	// after 5 seconds the heartbeat stops
	timeout := time.After(5 * time.Second)

	for {

		select {

		case <-heartbeatChannel:
			fmt.Println("heartbeat Received !")

		case <-timeout:
			fmt.Println("timeout!! heartbeat Stopped!")

			close(stop)

			return
		}
	}
}
