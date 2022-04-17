package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int
}

type SuperUser struct {
	user    User
	isSuper bool
}

func main() {
	user := User{name: "John", lastName: "Travolta", age: 44}
	sUser := SuperUser{user: user, isSuper: true}

	fmt.Printf("user: %v\nsUser: %v", user, sUser)
}
