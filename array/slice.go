package main

import "fmt"

func main() {

	var array []string = []string{"ali", "reza", "hosein", "max", "john", "mohammad", "hasan", "akbar"}

	slice := array[2:]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	array[2] = "alireza122"

	slice = append(slice, "majid")

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(cap(array))

	fmt.Println(array)

}
