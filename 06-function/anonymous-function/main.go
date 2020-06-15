package main

import "fmt"

func main() {
	func() { // this is anonymous function
		fmt.Println("hello this is anonymous function")
	}() // this is syntax for calling anonymous function

	func(x int) { // this is anonymous function
		fmt.Println("hello this is anonymous function with value : ", x)
	}(42) // this is syntax for calling anonymous function with passing value
}
