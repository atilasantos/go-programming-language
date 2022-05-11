package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Character struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var characters []Character

	err := json.Unmarshal([]byte(s), &characters)
	if err != nil {
		log.Fatal(err)
		return
	}

	for key, value := range characters {
		fmt.Printf("To key: %v we have the following attr:\n name: %v %v\n age: %v\n sayings: %v\n", key, value.First, value.Last, value.Age, value.Sayings)
	}

}
