package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println("function foo called")
	// foo(x) 		   // if we pass directly this variable it's will thrown error, bcs foo accept variadic type int not array of int
	foo(x...)          // instead we do like this or
	foo(1, 2, 3, 4, 5) // do this

	fmt.Println()

	fmt.Println("function boo called")
	boo("dhuki", x...)          // if we have multiple argument to pass and there is variadic type, we should make variadic type final argument
	boo("dhuki", 1, 2, 3, 4, 5) // like this put variadic type at the last of argument
}

func foo(x ...int) {
	fmt.Println(x)
}

func boo(a string, x ...int) {
	fmt.Println(a, x)
}
