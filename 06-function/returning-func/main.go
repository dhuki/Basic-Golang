package main

import "fmt"

func main() {

	// call this way
	fmt.Println("returning a func")
	x := bar()       // assign function to variable
	fmt.Println(x()) // call its function
	// or this way
	fmt.Println(bar()())
	fmt.Println()

	// another example func in go
	fmt.Println("returning a value with another style")
	fmt.Println(foo())
	fmt.Println(woo())
	fmt.Println()
}

func bar() func() int {
	return func() int { // return func of int
		return 123
	}
}

func foo() (s string) {
	s = "dhuki dwi"
	return
}

func woo() (s string, i int) {
	s = "dhuki dwi"
	i = 23
	return
}
