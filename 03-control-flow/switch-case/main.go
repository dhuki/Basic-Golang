package main

import "fmt"

func main() {

	// switch case with condition
	switch {
	case false:
		fmt.Println("this not going print")
	case 1 == 1:
		fmt.Println("this is going to print 1")
		fallthrough // this command will get into another condition below (but only one condition) and do something regardless if it's true or not
	case 2 == 2:
		fmt.Println("is this going to print 2 ?")
		fallthrough // remember this command
	case 3 == 4:
		fmt.Println("this is going to print 3, even though it false condition")
	case 4 == 4:
		fmt.Println("this is not going print because there is not fallthrough command on above condition")
	}

	fmt.Println()

	// switch case with value
	switch "dhuki" {
	case "dwi", "dhuki": // we can do this too with multiple value
		fmt.Println("this is not going to print")
	case "dhukid": // this will getting error if value statement is same "duplicate case "dhuki" in switch"
		fmt.Println("this is going to print")
	default:
		fmt.Println("not going meet all condition above")
	}

	fmt.Println()

	// switch case with type
	var x interface{}
	x = 12
	switch x.(type) { // it kinda like assertion
	case int: // we can do this too with multiple value
		fmt.Println("this is int type")
	case string: // this will getting error if value statement is same "duplicate case "dhuki" in switch"
		fmt.Println("this is string type")
	default:
		fmt.Println("this is someting else")
	}
}
