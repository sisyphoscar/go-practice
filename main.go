package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["First"] = "Oscar"
	myMap["Second"] = "Sisyphoscar"

	fmt.Println(myMap)
}
