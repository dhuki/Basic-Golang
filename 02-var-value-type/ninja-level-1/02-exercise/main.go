package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\t%v\t%v\t", x, y, z) // will print default value from parameter
}
