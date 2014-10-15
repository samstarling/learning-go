package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	p := Person{"Sam", "Starling"}
	q := &p
	q.FirstName = "Samuel"
	fmt.Println(q)
}
