package main

import "fmt"

type Person struct {
	firstName              string
	lastName               string
	favoriteIceCreamFlavor []string
}

func main() {
	person1 := Person{
		firstName: "Átila",
		lastName:  "Santos",
		favoriteIceCreamFlavor: []string{
			"Creme",
			"Chocolate",
			"Morango",
		},
	}

	person2 := Person{
		firstName: "Juanito",
		lastName:  "Cabalero",
		favoriteIceCreamFlavor: []string{
			"Limão",
			"Maracujá",
			"Marijuana",
		},
	}

	persons := make(map[string]Person)
	persons[person1.firstName] = person1
	persons[person2.firstName] = person2

	for key, value := range persons {
		fmt.Printf("%v information:\n", key)
		fmt.Printf("last_name: %v\n", value.lastName)
		fmt.Println("favorite_flavors:")
		for index, item := range value.favoriteIceCreamFlavor {
			fmt.Printf("\t%v - %v\n", index, item)
		}
	}

}
