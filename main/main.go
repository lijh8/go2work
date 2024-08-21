package main

import "iter"

// func Countdown(v int) func(func(int) bool) { // without iter package
func Countdown(v int) iter.Seq[int] { // with iter package
	return func(yield func(int) bool) {
		for i := v; i >= 1; i-- {
			if !yield(i) {
				println("yield: ", i)
				break
			}
			println("func: ", i)
		}
	}
}

func main() {
	n := 5
	for x := range Countdown(n) {
		println(x)
		if x == 3 {
			return // this make the yield return false and exit early
		}
	}
}
