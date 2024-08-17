package main

import (
	"fmt"
)

func main() {
	{
		a := [3]int{1, 2, 3}
		b := a[:] // ok, slice
		b[0] = 100
		fmt.Println(a, b, a[:])
	}

	{
		// string is immutable,
		// needs a type conversion explicitly
		// converting string to slice makes a copy
		a := "abc"
		b := []rune(a) // ok, slice,
		b[0] = 'A'
		fmt.Println(a, string(b), string(a[:]))
	}
}
