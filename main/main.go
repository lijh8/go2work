package main

import (
	"tuple2"
)

func main() {
	s1 := []any{1, "hello", 3.14, true}
	s2 := []any{1, "hello", 3.14, true}
	b, ok := tuple2.Cmp(s1, s2)
	println(b, ok)

}
