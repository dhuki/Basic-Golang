package main

import (
	"fmt"
)

func main() {
	// bi-directional channel
	c := make(chan int)

	// send channel
	go foo(c)
	// receive channel
	bar(c)
}

// since we can pass bi-directional to directional channel we can do this
// for more readble code
// send channel
func foo(c chan<- int) {
	c <- 42
}

// receive channel
func bar(c <-chan int) {
	fmt.Println(<-c)
}
