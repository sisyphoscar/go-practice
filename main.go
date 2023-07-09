package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthMonth  time.Month
}

func main() {
	me := User{
		FirstName:   "Oscar",
		LastName:    "Chiu",
		PhoneNumber: "0912345678",
		Age:         23,
		BirthMonth:  7,
	}

	myMap := make(map[string]User)

	myMap["me"] = me

	fmt.Println(myMap["me"].FirstName)
}
