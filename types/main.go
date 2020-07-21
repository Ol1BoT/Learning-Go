package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) details() string {
	return fmt.Sprintf("Name: %s %s, age: %d", p.firstName, p.lastName, p.age)
}

func main() {

	//structs
	ollie := person{
		"Ollie",
		"Mignot",
		31,
	}

	fmt.Println(ollie)
	fmt.Println(ollie.details())

	// arrays
	// arrays can't be mutable, and the size needs to be defined
	var arr = [5]int32{1, 2, 3, 4, 5}

	// fmt.Println(len(arr))
	fmt.Println(arr)

	//iter over a arr
	// for index, val := range arr {
	// 	fmt.Println(index, ":", val)
	// }

	//Slice
	//Slices are like arrays, but they can be appended and don't have a fixed len or cap
	var slice = []string{"Hello", "World"}

	slice = append(slice, "Another")
	// fmt.Println(slice)
	//iter over a slice, same as array
	for _, val := range slice {
		fmt.Println(val)
	}

	//string
	str := "Ollie"

	//iter over string
	for _, char := range str {
		fmt.Println(string(char))
	}
}
