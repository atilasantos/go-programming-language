package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	str := `[{"Name": "Pedro", "Age": 32},{"Name": "Carlo", "Age": 45},{"Name": "Atila", "Age": 10},{"Name": "Aaron", "Age": 5}]`
	bs := []byte(str)

	var people []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before sorting..")
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })

	fmt.Println("After sorting..")
	fmt.Println(people)
}
