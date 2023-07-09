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
	user := User{
		FirstName:   "Oscar",
		LastName:    "Chiu",
		PhoneNumber: "0912345678",
		Age:         23,
		BirthMonth:  7,
	}

	user.printFirstName()
}

func (user *User) printFirstName() {
	fmt.Println(user.FirstName)
}
