package main

import "fmt"

func main() {
	age := 31
	birth_date := 2022 - age
	for birth_date <= 2022 {
		fmt.Println(birth_date)
		birth_date++
	}
}
