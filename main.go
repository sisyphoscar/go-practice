package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []string
	numbers := []int{3, 2, 1}

	mySlice = append(mySlice, "Oscar")
	mySlice = append(mySlice, "Sisyphoscar")
	numbers = append(numbers, 4)

	sort.Ints(numbers)

	fmt.Println(mySlice)
	fmt.Println(numbers)
}
