package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

func (p Person) Talk() string {
	return fmt.Sprintf("Hello, %v!\n", p.name)
}

func main() {
	person := Person{name: "Átila", age: 31, email: "atila.romao@hotmail.com"}
	fmt.Println(person.Talk())
}
