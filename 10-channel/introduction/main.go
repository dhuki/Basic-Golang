package main

func main() {
	// hint : to use channel for conccurency in go
	// Do not comunicate by sharing memomy / sharing variable, instead, share memory by comunication
	// this is example bi-directional channel

	// error deadlock
	// c := make(chan int)
	//
	// c <- 42 // need another channel to pull value from main goroutine here
	// when there is no another channel, it will stack and get deadlock
	// bcs it searching for another goroutine to pull a value
	// fmt.Println(<-c)

	// solution 1
	// go func() {
	// 	c <- 2
	// }()
	// fmt.Println(<-c)

	// solution 2 using buffered channel
	// it bcs we allowing to accept one go routine and pull value at the same goroutine
	// c := make(chan int, 1)

	// c <- 2

	// fmt.Println(<-c)
}
