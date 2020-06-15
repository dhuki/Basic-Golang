package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"` // use tags to change the name of json
	Last  string `json:"last"`  // we should make field can exported so uppercase at first char, if its not will gettin empty json
	Age   int    `json:"age"`   // we should make field can exported so uppercase at first char, if its not will gettin empty json
}

func main() {
	p1 := person{
		First: "dhuki",
		Last:  "dwi",
		Age:   23,
	}

	p2 := person{
		First: "ujang",
		Last:  "bedil",
		Age:   23,
	}

	// convert to string
	people := []person{p1, p2}
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bs)) // coversion byte of array to string

	// convert to json
	s := `[{"first":"dhuki","last":"dwi","age":23},{"first":"ujang","last":"bedil","age":23}]`
	bs2 := []byte(s)

	// people := {}person[]
	var people2 []person                // using this bcs more readable
	err = json.Unmarshal(bs2, &people2) // passing pointer bcs it will change value of people
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("all data")

	for _, v := range people2 {
		fmt.Println(v)
	}
}
