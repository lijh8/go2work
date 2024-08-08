package main

import (
	"tuple2"
)

func main() {
	b := []any{1, "hello", 3.1, true}
	a := []any{1, "hello", nil, true}
	c, ok := tuple2.Cmp(a, b)
	println(c, ok)

}
