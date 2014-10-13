package main

import "fmt"

func main() {

	// Declare a variable
	fmt.Println("Declaring variables")
	var i int
	var s string
	fmt.Println(i, s)

	// Variables with initializers
	fmt.Println("Variables with initializers")
	var j int = 3
	var t string = "foo"
	fmt.Println(j, t)

	// Multiple variables
	var a, b int = 1, 2
	fmt.Println(a, b)

	// Short assignment
	fmt.Println("Short assignment")
	x := 3
	fmt.Println(x)

}
