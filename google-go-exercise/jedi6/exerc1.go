package main

import "fmt"

func main() {

	fnFoo := foo()
	fnbarint, fnbarstring := bar()

	fmt.Println(fnFoo)
	fmt.Println(fnbarint, fnbarstring)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "Rule the world"
}
