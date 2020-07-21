package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	age       int8
}

func main() {

	ollie := person{
		"Ollie",
		"Mignot",
		31,
	}

	jn, err := json.Marshal(ollie)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jn))
}
