package main

import "fmt"

func main() {
	integers := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice1 := append(integers[2:7])
	slice2 := append(integers[7:])
	slice3 := append(integers[4:9])
	slice4 := append(integers[3:8])

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
