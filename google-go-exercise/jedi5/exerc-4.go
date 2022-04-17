package main

import "fmt"

type Human interface {
	Eat()
	Speak()
	Walk()
}

type Person struct {
	name  string
	age   int
	email string
}

type Chinese struct {
	person Person
	coin   string
}

type Brazilian struct {
	person Person
	coin   string
}

func (c Chinese) Eat() string {
	return fmt.Sprintf("barata")
}

func (b Brazilian) Eat() string {
	return fmt.Sprintf("brigadeiro")
}

func (p Person) Talk() string {
	return fmt.Sprintf("Hello, %v!\n", p.name)
}

func main() {
	person := Person{name: "Átila", age: 31, email: "atila.romao@hotmail.com"}
	fmt.Println(person.Talk())

	person2 := Brazilian{person: person, coin: "Real"}
	person3 := Chinese{person: person, coin: "Ien"}

	fmt.Println(person2.Eat())
	fmt.Println(person3.Eat())
}
