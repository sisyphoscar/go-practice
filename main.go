package main

import "fmt"

func main() {
	animal := "cat"

	if animal == "human" {
		fmt.Println("human")
	}

	switch animal {
	case "cat":
		fmt.Println("cat")
	case "dog":
		fmt.Println("dog")
	default:
		fmt.Println("not matched")
	}
}
