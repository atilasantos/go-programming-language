package main

import "fmt"

type atila int

var x atila

func main() {

	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)

}
