package main

import (
	"fmt"
	"math/cmplx"
)

// Can define many vars like this
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Here are some datatypes
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// T(v) converts value v to type T
	var i int = 42
	var s string = string(42)
	fmt.Println(i, s)
}
