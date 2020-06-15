package main

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person        // name variable should same with types if we want use embedded struct, it's call anonymous field (or it's call inheritance in OOP)
	name   string // it's call identifier list and its type
	kill   bool
}

func main() {
	sa := secretAgent{
		person: person{ // name variable in person type should have same name with the type
			name: "dhuki", // if we had same name variable with inner class defined will (sa.person.name)
			age:  23,
		},
		name: "dwi",
		kill: true,
	}

	fmt.Println(sa.person.name, sa.name, sa.age, sa.kill)
}
