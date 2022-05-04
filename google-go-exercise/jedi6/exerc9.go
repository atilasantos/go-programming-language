package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	age, _ := strconv.Atoi(os.Args[2]) //geting the second argument passed withing the code execution
	greeting := giveGreeting(greetings, os.Args[1], age)

	fmt.Println(greeting)
}

func greetings(name string, age int) string {
	return fmt.Sprintf(`Hello, %v i'm %v old`, name, age)
}

func giveGreeting(f func(s string, i int) string, name string, age int) string {
	return f(name, age)
}
