package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7)
}

func foo(x ...int) { // this parameter will has type of int slice ([]int)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
