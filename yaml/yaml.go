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

type config struct {
	Base     base
	Addition struct {
		Firstname string
		Lastname  string
	}
}

func main() {

	file, err := ioutil.ReadFile("example.yml")
	if err != nil {
		panic(err)
	}

	var config config

	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	fmt.Println(config)

	fmt.Printf("%+v\n", config)
	fmt.Printf("%v\n", config)

	fmt.Println(config.Base.URL)

}
