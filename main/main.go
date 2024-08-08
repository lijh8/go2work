package main

import (
	"tuple2"
)

func main() {
	a := []any{1, "hello", 3.14, true}
	b := []any{1, "hello", 3.14, true}

	c, ok := tuple2.Cmp(a, b)
	println(c, ok)

}
