package main

import (
	"fmt"
)

func main() {
	{
		a := [3]int{1, 2, 3}
		b := a[:] // ok, slice
		b[0] = 100
		fmt.Println(a, b)
	}

	{
		// string is immutable like the string literal in c language.
		// it needs a type conversion to slice explicitly.
		// char []a = "hello";
		a := "abc"
		b := []rune(a) // ok, slice,
		b[0] = 'A'
		fmt.Println(a, string(b))

	}
}
