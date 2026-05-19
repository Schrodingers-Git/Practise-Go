package main;
/*
	Livelock Implementation in Go

	a single pathway,Alice and Bob expect each other to move first,making no progress.

	Deadlock Vs Livelock 

	=> Deadlocks share some common variable/ object    vs   => Livelocks dont share any common variable
	=> Deadlocks can be prevented by Locks(Mutex)      vs   => Logical Implementation has error,that should be corrcected.
*/


import (
	"fmt"
	"time"
)

func worker(myName string, myStatus bool, opponentStatus bool) {

	for opponentStatus{

		if myName == "Alice" {
			fmt.Println("Alice is waiting for Bob to move")
		} else {
			fmt.Println("Bob is waiting for Alice to move")
		}

		time.Sleep(1 * time.Second)

	}
		fmt.Println(myName, "is Moving")
}

func main() {

	aliceStatus := true
	bobStatus := true

	go worker("Alice", aliceStatus, bobStatus)
	go worker("Bob", bobStatus, aliceStatus)

	time.Sleep(5 * time.Second)
}
