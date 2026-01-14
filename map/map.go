package main

import "fmt"

func main() {

	var names map[int]string = map[int]string{}

	names[1250820709] = "ali"
	names[1250822709] = "reza"
	names[1250820409] = "hosein"

	fmt.Println(names)

}
