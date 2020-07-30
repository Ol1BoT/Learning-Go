package main

import "fmt"

func main() {
	//anonymous function
	func() {
		fmt.Println("Testing")
	}()
	//anonymous function with param
	func(x int) {
		fmt.Println("The param is:", x)
	}(42)

	//func expression - feels similar to JavaScript
	//func is a type

	f := func() {
		fmt.Println("This is a func expression")
	}

	f()

}
