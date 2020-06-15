package main

import "fmt"

func main() {
	value := 42

	// \n , \t , many more is scaped character
	// %v is default value, by its variable
	fmt.Printf("this is formatted string with value of variable : %v\n", value)
	// %T is type of value, by its variable
	fmt.Printf("this is formatted string with type of variable : %T\n", value)
	// %T is binary (0/1) of value, by its variable (string doesn't have this)
	fmt.Printf("this is formatted string with byte of variable : %b\n", value)
	// %x is hex of value without 0x, by its variable
	fmt.Printf("this is formatted string with hex of variable : %x\n", value)
	// %x is hex of value with 0x, by its variable
	fmt.Printf("this is formatted string with hex of variable : %#x\n", value)
	// using multiple argument
	fmt.Printf("this is formatted string with hex of variable : %v\t%T\t%b\n", value, value, value)

	// Sprintf it will making string value
	s := fmt.Sprintf("this is formatted string with hex of variable : %v\t%T\t%b", value, value, value)
	fmt.Println(s)
}
