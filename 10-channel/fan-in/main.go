package main

import (
	"fmt"
	"sync"
)

func main() {
	// fan-in basically means thet we pull from multiple channel to one channel
	// just like multiplexer
	// design pattern concurreny

	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send channel
	go foo(eve, odd)
	// receive
	go bar(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

// this is fan-in that we used
func bar(eve, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range eve {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func foo(eve, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}
