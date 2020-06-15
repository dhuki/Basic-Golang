package main

import "fmt"

// underlying type of hello is int
type hello int

var something hello

func main() {
	fmt.Printf("%v\n", something)
	fmt.Printf("%T\n", something)

	something = 42

	fmt.Printf("%v\n", something)
}
