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

// yield returning false and early exit
func main() {
	n := 5
	value := 3
	found := false
	for x := range Countdown(n) {
		if x == value {
			found = true
			break // this make the yield return false and exit early
		}
	}
	println(found)
}
