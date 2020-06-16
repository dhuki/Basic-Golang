package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	done := make(chan int)

	// send channel
	go foo(eve, odd, done)

	// receive
	bar(eve, odd, done)
}

func bar(eve, odd, done <-chan int) {
	for { // return forever bcs foo method send many value and we have to catch it all
		select { // select is just like switch case but usually use for channel
		case v := <-eve:
			fmt.Println("this is even:\t", v)
		case v := <-odd:
			fmt.Println("this is odd:\t", v)
		// case v := <-done:
		// 	fmt.Println("and done:\t", v)
		// 	return
		case i, ok := <-done: // let's we are using idiomatic golang (, ok)
			if !ok {
				fmt.Println("done channel is closed and value is : ", i) // closed channel has zero value in it
				return
			}
		}
	}
}

func foo(eve, odd, done chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	// done <- 1
	close(done) // it will has 2 value
}
