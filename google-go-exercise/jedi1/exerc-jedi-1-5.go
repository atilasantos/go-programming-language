package main

import "fmt"

type atila int

var x atila

var y int

func main() {

	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)

	y = int(x)

	fmt.Printf("y value %v\ny type %T", y, y)

}
