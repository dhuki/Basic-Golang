package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// receiver is value not pointer
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// pointer is use when you had large of data and you have to pass it to your apps
	// if you need to changes something in certain location addresss

	a := 42
	fmt.Println(a)
	fmt.Println(&a) // ampersand operator (&) give an address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // type pointer of int

	fmt.Println()
	b := &a
	fmt.Println(*b) // dereference the address and get value
	fmt.Println(*&a)
	fmt.Println()

	// 1. use if we need to changes
	x := 12
	fmt.Println(x)
	// we need change x to something
	foo(&x)
	fmt.Println(x)
	fmt.Println()

	fmt.Println("receiver/method set")
	// receiver 	value
	// (t T)		T and *T
	// (t *T)		*T
	c := circle{5}
	// receiver value can accept pointer or not pointer value
	// again the same thing we use pointer receiver if we have to change
	// something in certain location
	info(c)
	info(&c)
	changeMe(&c) // it change value by pass reference, again we use pointer receiver if we have to change
	info(c)
	info(&c)
}

func changeMe(c *circle) {
	c.radius = 15    // we can change like this for struct
	(*c).radius = 15 // or maybe like this one
}

func foo(y *int) {
	*y = 13
}
