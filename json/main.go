package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int8   `json:"age"` //originally had this named as age lower case, and the value wasn't being marshalled to json
}

func main() {

	ollie := person{
		"Ollie",
		"Mignot",
		31,
	}

	jn := toJSON(&ollie)

	fmt.Println(ollie)

	fmt.Println(string(jn))

}

func toJSON(p *person) []byte {
	jn, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(jn))

	return jn
}
