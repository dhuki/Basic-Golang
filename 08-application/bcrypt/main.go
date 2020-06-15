package main

import (
	"fmt"
	// should import bcrypt here
)

func main() {
	// excellent way to store password bcs it cannot decript but you can match it
	password := "123456"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost) // minconst actually just how hard password will be encrypt
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(password)
	fmt.Println(string(bs))
	fmt.Println()

	// matching password bcrypt
	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("You're login")
}
