package main

import (
	"fmt"
)

func main() {

	var users [3]string = [3]string{"hosein", "max", "john"}

	usersCopy := &users

	changeValue(usersCopy)

	fmt.Println(usersCopy)
	fmt.Println(users)

}

func changeValue(users *[3]string) {
	users[0] = "Alice"
}
