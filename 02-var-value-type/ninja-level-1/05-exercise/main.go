package main

import "fmt"

// underlying type of hello is int
type hello int

var something hello

var x int

func main() {
	fmt.Printf("%v\n", something)
	fmt.Printf("%T\n", something)

	something = 42

	fmt.Printf("%v\n", something)

	x = int(something)
	fmt.Println(x)
	fmt.Printf("%T\t", x)
}
