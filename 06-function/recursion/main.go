package main

import "fmt"

var total = 1

func main() {
	// making factorial
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factorial2(4))
	fmt.Println()

	// fibonaci
	fmt.Println(1 + 1 + 2 + 3 + 5)
	fmt.Println(fibonaci(4))
}

func factorial(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}

func factorial2(value int) int {
	sum := 1
	for ; value > 0; value-- {
		sum *= value
	}
	return sum
}

func fibonaci(length int) int {
	sum := 0
	next := 1
	for i := 0; i < length; i++ {
		result := sum + next
		next = sum
		sum += result
	}
	return sum
}
