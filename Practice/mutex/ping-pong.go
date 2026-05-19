package main


//NOTE: A. Ping Pong Implementation => Go Routines , Channels , Mutex Locks , WaitGroups

//NOTE:Bad Implementation (just tryin the concepts)
/*

	Routines  -> Alice,Bob 
	Channels  -> Alice,Bob,Done
	Mutex Lock for the done_counter 
	WaitGroups to wait till both the go Routines are completed.
	
	=> But here(in this Implementation) since only one Routine ends the program,(i.e) 
		if Alice Routine goes to the end condition(cntr>done_counter),
		then it gets ended,but the Bob Routine waits by listening to the channel 
		and the WaitGroups also wait for the Bob to close,So it results in a Deadlock / infinite time execution.
	
*/
/*
import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex

func main() {

	fmt.Println("Ping Pong Game - Go Routine n Channel Practise")
	
	//done_counter 
	var done_count int
	fmt.Scan(&done_count)
	
	wg := sync.WaitGroup{}

	ch_a := make(chan string)
	ch_b := make(chan string)
	ch_done := make(chan string)

	// Alice
	wg.Add(1)
	
	go func() {
		for {

			msg := <-ch_a

			if msg == "Bob Received,Sent back to Alice - Pong" {

				mu.Lock()

				counter++

				fmt.Println("Alice Counter:", counter)

				if counter >= done_count {
					mu.Unlock()
					ch_done <- "Game Over !"
					wg.Done()
					wg.Done()
					return
				}

				mu.Unlock()

				time.Sleep(time.Second)

				ch_b <- "Alice Received,Sent back to Bob - Ping"
			}
		}
	}()

	// Bob
	wg.Add(1)

	go func() {

		for {

			msg := <-ch_b

			if msg == "Alice Received,Sent back to Bob - Ping" {

				mu.Lock()

				counter++

				fmt.Println("Bob Counter:", counter)

				if counter >= done_count {
					mu.Unlock()
					ch_done <- "Game Over !"
					return
				}

				mu.Unlock()

				time.Sleep(time.Second)

				ch_a <- "Bob Received,Sent back to Alice - Pong"
			}
		}
	}()
		

	// Start Game
	ch_a <- "Bob Received,Sent back to Alice - Pong"
	
	wg.Wait()

	// Proper synchronization
	for{
		msg := <-ch_done
		fmt.Println(msg)
	}
}

*/




//NOTE:Clean Implementation

/*
The above Implementation uses 2 channels to communicate between the Alice and Bob ,
Here we use a single channel -> for the ball,
and simulate the alternate turn behaviours using the sleep methods in the routines.
and ensure both the routines are ended using the WaitGroups
*/


import (
	"fmt"
	"time"
	"sync"
)

/*
Why struct here? 
Couples multiple things like lastHitBy,counter,other details etc into one.
if we were to use the global counter variable ,then i need to care about the mutxes ,
But now we use the channel , pointer of the struct and sleep() so both does'nt access it at the same time.

And why send pointer of the struct? 
	Else u send the copy of struct not the actual struct instance u created.
*/

type Ball struct {
	counter   int
	lastHitBy string
	maxHits   int
}

//The "Worker"
func player_hit(name string, ch chan *Ball, wg *sync.WaitGroup) {
	
	defer wg.Done()

	for b := range ch {
		b.counter++
		b.lastHitBy = name

		fmt.Printf("Round %d: %s hit the ball!\n", b.counter, name)

		if b.counter >= b.maxHits {
			fmt.Printf("--- %s reached the limit. Game Over! ---\n", name)
			close(ch) // Closing the channel kills the 'range' loop for EVERYONE
			return
		}

		time.Sleep(500 * time.Millisecond)

		ch <- b
	}
}

func main() {
	var inputHits int
	fmt.Print("Enter max hits for the game: ")
	fmt.Scan(&inputHits)

	// Create the channel for the ball
	ch := make(chan *Ball)

	// WaitGroup for synchronization
	wg := sync.WaitGroup{}

	// CRITICAL: Add(2) happens in the parent (main) before the goroutines start.
	wg.Add(2)

	// Start Alice and Bob Routines
	// Note: We pass the pointer to the WaitGroup so they can call Done().
	go player_hit("Alice", ch, &wg)
	go player_hit("Bob", ch, &wg)

	// Start the Game
	ch <- &Ball{
		counter: 0,
		maxHits: inputHits,
	}

	// Block main until both players return
	wg.Wait()
}
