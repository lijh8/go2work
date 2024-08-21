package main

import "iter"

// func Countdown(v int) func(func(int) bool) { // without iter package
func Countdown(v int) iter.Seq[int] { // with iter package
	return func(yield func(int) bool) {
		for i := v; i >= 0; i-- {
			if !yield(i) {
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
}

// output:
// 3
// 2
// 1
// 0
