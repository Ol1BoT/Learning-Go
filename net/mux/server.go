package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/test/", t)

	http.ListenAndServe(":8080", nil)

}

func t(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Test Route")
}
