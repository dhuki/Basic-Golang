package main

import "fmt"

func main() {
	p := struct { // it's call anonymous struct
		first string
		last  string
		age   int
	}{
		first: "dhuki",
		last:  "dwi",
		age:   23,
	}

	fmt.Println(p)
}
