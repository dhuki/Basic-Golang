package main

import "fmt"

// primitive type
var a int

// making own type
type sosis int

var b sosis

func main() {
	a = 10
	b = 11
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b) // it will print type of b is main.sosis

	// a = b // it will getting error since a is type of integer and b is type of main.sosis
	// to solve this we can use conversion
	a = int(b)
	fmt.Printf("%v\t%T\n", a, a)

}
