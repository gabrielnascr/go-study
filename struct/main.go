package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

func main() {
	var users []User

	users = append(users, User{
		FirstName: "Gabriel",
		LastName:  "Nascimento",
		Age:       15,
		Email:     "taltal@gmail.com",
	})

	users = append(users, User{
		FirstName: "José",
		LastName:  "Pereira",
		Age:       57,
		Email:     "seuze@gmail.com",
	})

	for i, user := range users {
		fmt.Printf("%d %s %s %d %s\n", i+1, user.FirstName, user.LastName, user.Age, user.Email)
	}
}
