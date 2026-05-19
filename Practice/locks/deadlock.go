package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    lock1 sync.Mutex
    lock2 sync.Mutex
)

func routine1() {
    lock1.Lock()
    fmt.Println("Routine 1: acquired lock1")
    time.Sleep(100 * time.Millisecond)
    
    lock2.Lock() // DEADLOCK: waiting for routine2 to release lock2
    fmt.Println("Routine 1: acquired lock2")
    lock2.Unlock()
    lock1.Unlock()
}

func routine2() {
    lock2.Lock()
    fmt.Println("Routine 2: acquired lock2")
    time.Sleep(100 * time.Millisecond)
    
    lock1.Lock() // DEADLOCK: waiting for routine1 to release lock1
    fmt.Println("Routine 2: acquired lock1")
    lock1.Unlock()
    lock2.Unlock()
}

func main() {
    go routine1()
    go routine2()
    
    time.Sleep(2 * time.Second)
    fmt.Println("Main exits (deadlock occurred)")
}

//NOTE:Fixing the deadlock by ensuring both routines acquire locks in the same order (e.g., lock1 -> lock2) can prevent the deadlock.
/*
func routine1() {
    lock1.Lock()
    lock2.Lock() // Always lock1 -> lock2
    fmt.Println("Routine 1: work done")
    lock2.Unlock()
    lock1.Unlock()
}

func routine2() {
    lock1.Lock() // Same order: lock1 -> lock2
    lock2.Lock()
    fmt.Println("Routine 2: work done")
    lock2.Unlock()
    lock1.Unlock()
}
*/
