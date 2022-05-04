package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeMe(p *Person) {
	p.age = 42
	p.name = `Carlo`

}

func main() {

	person := Person{name: "Cremilda", age: 54}
	fmt.Println(person)
	changeMe(&person)
	fmt.Println(person)

}
