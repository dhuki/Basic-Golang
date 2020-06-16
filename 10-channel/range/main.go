package main

import "fmt"

func main() {
	// bi-directional channel
	c := make(chan int)

	// send channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i // it will send 100 channel
		}
		close(c) // make sure close the channel *
	}()

	// receive channel using range
	// make sure close channel when sending channel *
	// without close channel it will keep wait and making error deadlock bcs we using range
	for v := range c { // it's gonna read that one by one
		fmt.Println(v)
	}
}
