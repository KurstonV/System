package main

import (
	"fmt"
)

// A synchronous operation - sending & receiving has to 
// happen before we can proceed
func main() {
    ch := make(chan int) // Unbuffered channel
    go func() {
       ch <- 8 // blocks until main goroutine receives
    }()
    value := <-ch // Receives value and unblocks goroutine
    fmt.Println("Received:", value)
}
