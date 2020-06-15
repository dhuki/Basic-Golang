package main

import "fmt"

func main() {
	a := incrementor()
	// wanna see a magic ??
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func incrementor() func() int {
	var x int           // zero value of int is 0
	return func() int { // this is closure
		x++
		return x
	}
}
