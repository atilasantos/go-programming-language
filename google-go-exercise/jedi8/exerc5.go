package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func iterateAndPrint(users []user) {
	fmt.Println("====================")
	for _, user := range users {
		fmt.Printf("%v %v\n", user.First, user.Last)
		sort.Strings(user.Sayings)
		for _, v := range user.Sayings {
			fmt.Printf("\t\n%v", v)
		}
		fmt.Println()
		fmt.Println("++++++++++++++++++++")
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Not sorted")
	fmt.Println(users)

	// by age
	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	iterateAndPrint(users)

	// by last
	sort.Slice(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	iterateAndPrint(users)

}
