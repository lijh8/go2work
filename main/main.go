package main

import (
	"fmt"
	"tuple2"
)

func main() {
	a := []any{"abc", 123, 3.14, []any{"abc", 1234, "aaa"}}
	b := []any{"abc", 123, 3.14, []any{"abc", 123, 3.14}}
	c, ok := tuple2.Cmp(a, b)
	fmt.Println(c, ok)
}
