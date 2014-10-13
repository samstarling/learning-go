package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(swap("a", "b"))
	fmt.Println(two_strings())

	foo, bar := named_results()
	fmt.Println(foo, bar)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func two_strings() (string, string) {
	return "Sam", "Starling"
}

func named_results() (foo, bar string) {
	return "foo", "bar"
}
