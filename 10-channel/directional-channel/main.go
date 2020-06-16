package main

import "fmt"

func main() {
	// example 1
	c := make(chan<- int, 1) // this will got error bcs this channel is send-only cannot be read / receive
	c <- 2
	fmt.Println(<-c) // this is read / receive from channel

	// example 2
	ch := make(<-chan int, 1) // this will got error otherwise this channel is read-only / receive-only cannot send something
	ch <- 2                   // this is send to the channel
	fmt.Println(<-ch)         // this is read / receive from channel

}
