package main

import "fmt"

func main() {
	name := "Oscar"

	fmt.Println(SayHelloToUser(name))

	changeUserNameByPointer(&name, "Sisyphoscar")

	fmt.Println(SayHelloToUser(name))
}

func SayHelloToUser(user string) string {
	return "Hello " + user + "!"
}

func changeUserNameByPointer(name *string, newName string) {
	*name = newName
}
