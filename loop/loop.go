package main

import "fmt"

func main() {

	var users []user = []user{
		{code: 1, name: "hosein", email: "hosein@gmail.com", phone: "09905491724"},
		{code: 2, name: "reza", email: "reza@gmail.com", phone: "09905891754"},
		{code: 3, name: "mohammad", email: "mohammad@gmail.com", phone: "09905491724"},
		{code: 4, name: "max", email: "max@gmail.com", phone: "09905871724"},
		{code: 5, name: "john", email: "john@gmail.com", phone: "09903891724"},
		{code: 6, name: "david", email: "david@gmail.com", phone: "09705891724"},
		{code: 7, name: "trump", email: "trump@gmail.com", phone: "09305891724"},
	}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

}

type user struct {
	code  int
	name  string
	email string
	phone string
}
