package main

import "fmt"

// Stage 1: Generate numbers
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Square numbers
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Stage 3: Add 10 to numbers
func add10(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n + 10
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: gen -> sq -> add10
    for result := range add10(sq(gen(1, 2, 3, 4))) {
        fmt.Println(result)
    }
    // Output: 11, 14, 19, 26
}
