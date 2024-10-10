package main

import (
	"slices"
	"tuple2"
)

type S struct {
	name string
	num  int
}

func cmp(a, b S) int {
	x := []any{a.name, a.num}
	y := []any{b.name, b.num}
	z, _ := tuple2.Compare(x, y)
	return z
}

func main() {
	haystack := []S{
		{"aaa", 10},
		{"bbb", 20},
		{"ccc", 30},
		{"bbb", 20},
		{"bbb", 20},
	}
	needle := S{"bbb", 20}
	slices.SortFunc(haystack, cmp)
	lower, _ := slices.BinarySearchFunc(haystack, needle, cmp)
	upper := lower
	for upper != len(haystack) && cmp(needle, haystack[upper]) != -1 {
		upper++
	}
	LOG(lower, upper)
}
