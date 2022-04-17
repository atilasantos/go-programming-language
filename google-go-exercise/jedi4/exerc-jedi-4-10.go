package main

import "fmt"

func main() {

	person := make(map[string][]string)

	person["bond_james"] = []string{
		`Shaken, not stirred`, `Martinis`, `Women`,
	}

	person["moneypenny_miss"] = []string{
		`James Bond`, `Literature`, `Computer Science`,
	}

	person["no_dr"] = []string{
		`Being evil`, `Ice cream`, `Sunsets`,
	}

	for key, value := range person {
		fmt.Printf("This is the record for %v\n", key)
		for index, v := range value {
			fmt.Println("\t", index, v)
		}

	}

	person["Todd McLeod"] = []string{
		"Learn to Code Go on Udemy",
		"Part 1",
		"Page 45",
	}

	for key, value := range person {
		fmt.Println(key, value)
	}

	delete(person, "no_dr")

	for key, value := range person {
		fmt.Println(key, value)
	}

}
