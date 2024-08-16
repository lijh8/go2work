package main

import (
	"fmt"
	"tuple2"
)

func main() {
	a := []any{"abc", 123, 3.14}
	b := []any{"abc", 123, 3.14}
	c, ok := tuple2.Compare(a, b)
	fmt.Println(c, ok)
}
