package main

import "fmt"

func main() {

	// Generate
	var persons = make(map[string]Person)

	// Initial
	persons["1250828707"] = Person{Name: "Hosein", Family: "Tahan"}
	persons["1250810201"] = Person{Name: "Reza", Family: "Hoseini"}
	persons["1250230103"] = Person{Name: "Mohammad", Family: "Musk"}
	persons["1250243385"] = Person{Name: "Donald", Family: "Trump"}

    // Delete
	delete(persons,"1250243385")

	// Update
	persons["125023013"]=Person{Name: "updatedName",Family: "updatedFamily"}

	// Get
	person,ok :=persons["1250828707"]

	// Check
	if ok {
		fmt.Printf("Hello %s \n",person.Name)
	}else{ 
	   fmt.Println("Person Not Found")
	}
    
	fmt.Println(persons)

}

type Person struct {
	Name   string
	Family string
}