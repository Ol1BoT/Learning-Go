package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.google.com")
	errorChecker(err)
	defer resp.Body.Close()

	// fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	errorChecker(err)
	fmt.Println(string(body))

}

func errorChecker(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
