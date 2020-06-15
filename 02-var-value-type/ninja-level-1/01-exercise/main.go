package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	// single print statement
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

	// multiple print statement
	fmt.Printf("%v\t%v\t%v\t", x, y, z)
}
