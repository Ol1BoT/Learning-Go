package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type base struct {
	URL  string
	Port int
	User string
}

//Strings in the Yaml don't have to be in quotes

type config struct {
	Base     base
	Addition struct {
		Firstname string
		Lastname  string
		Likes     []string
	}
}

func main() {

	file, err := ioutil.ReadFile("example.yml")
	if err != nil {
		panic(err)
	}

	//slice of bytes cast into a string
	fmt.Println(string(file))

	var config config

	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	fmt.Println(config)

	fmt.Printf("%+v\n", config)
	fmt.Printf("%v\n", config)

	fmt.Println(config.Base.URL)

}
