package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	p := Person{"Sam", "Starling"}
	q := &p
	fmt.Println(q)
}
