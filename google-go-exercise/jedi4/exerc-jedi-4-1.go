package main

import "fmt"

func main() {
	integers := [5]int{10, 20, 30, 40, 50}

	for _, value := range integers {
		fmt.Printf("%v %T\n", value, value)
	}
}
