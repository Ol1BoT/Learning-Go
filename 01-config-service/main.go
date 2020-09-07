package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ol1bot/Learning-Go/01-config-service/domain"
)

func main() {
	fmt.Printf("This is the main package\n")

	config := domain.Config{}

	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = config.SetFromBytes(data)
	if err != nil {
		panic(err)
	}

	cfg, err := config.Get("addition")
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
