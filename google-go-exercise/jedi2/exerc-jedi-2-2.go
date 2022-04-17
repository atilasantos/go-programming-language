package main

import "fmt"

func main() {

	g := 10 == 20
	h := "true" <= "true"
	i := 1 >= 0
	j := "hi" != "Hi"
	k := "a" < "A"
	l := "a" > "A"

	fmt.Printf("g: %v\th: %v\t\ti: %v\nj: %v\t\tk: %v\tl: %v", g, h, i, j, k, l)
}
