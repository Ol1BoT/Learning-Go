package main

import (
	"fmt"
	"net/http"
)

type myHander int

func (h myHander) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is my Server")
}

func main() {
	var h myHander
	http.ListenAndServe(":8080", h)

}
