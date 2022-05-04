package main

import "fmt"

func main() {

	person := Person{first: "Átila", last: "Santos", age: 31}
	person.Speak()
}

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) Speak() {
	fmt.Printf("name: %v\tage: %v", p.first, p.age)
}
