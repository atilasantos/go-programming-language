package main

import "fmt"

func main() {
	sheet := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, v := range sheet {
		for _, va := range v {
			fmt.Println(va)
		}
	}
}
