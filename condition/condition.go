package main

import "fmt"

func main() {

	var age int

	fmt.Print("Please Enter Your Age : ")

	fmt.Scanln(&age)

	if age <= 10 {
		fmt.Print("You Are BABY")
	} else if age > 10 && age <= 15 {
		fmt.Print("You Are TEANAGER")
	}else if age > 15 && age <= 45 {
		fmt.Print("You Are YOUNG")
	}else if age > 45 && age <= 100 {
		fmt.Print("You Are OLD")
	}

}