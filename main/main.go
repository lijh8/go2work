package main

import "iter"

// func Countdown(v int) func(func(int) bool) { // without iter package
func Countdown(v int) iter.Seq[int] { // with iter package
	return func(yield func(int) bool) {
		for i := v; i >= 1; i-- {
			if !yield(i) {
				break
			}
		}
	}
}

// func DemoSeq2[T any](s []T) func(func(int, T) bool) { // without iter package
func DemoSeq2[T any](s []T) iter.Seq2[int, T] { // with iter package
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				break
			}
		}
	}
}

func main() {
	n := 3
	for x := range Countdown(n) {
		println(x)
	}

	s := []string{"aaa", "bbb", "ccc"}
	for i, v := range DemoSeq2(s) {
		println(i, v)
	}
}

// output:
// 3
// 2
// 1
// 0 aaa
// 1 bbb
// 2 ccc
