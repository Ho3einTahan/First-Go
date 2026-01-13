package main

import "fmt"

func main() {

	numbers := [10]int{10, 20, 30, 40}

	var search int = 30

	fmt.Println(numbers)

	for index, number := range numbers {
		if number == search {
			fmt.Printf("The Number %d Founded At %d", number, index)
			break
		}
	}

}
