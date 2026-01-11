package main

import "fmt"

func main() {

	var degree float32

	fmt.Print("Please Enter Weather degree : ")

	fmt.Scanln(&degree)

	switch {
	case degree > (-12) && degree < (-5):
		fmt.Print("COLD")
	case degree > (10) && degree < (30):
		fmt.Print("Middle")
	case degree > 30 && degree < 40:
		fmt.Print("Warm")
	case degree > 40 && degree < 55:
		fmt.Print("Hot")
	default:
		fmt.Print("Degree Is Out Of Range")
	}

}
