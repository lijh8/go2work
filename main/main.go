package main

import "iter"

// standard push iterator
func Countdown(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := v; i >= 1; i-- {
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

// https://golang.google.cn/blog/range-functions#pull-iterators ,

// Pull iterators are not supported directly by the for/range statement;
// It is straightforward to write an ordinary for statement that loops
// through a pull iterator.

// pull iterator
func CountdownPull(v int) func() (int, bool) {
	current := v
	return func() (int, bool) {
		if current > 0 {
			val := current
			current--
			return val, true
		}
		return 0, false
	}
}

func main() {
	n := 3
	// Create the pull iterator
	next := CountdownPull(n)

	// Use the pull iterator
	for {
		val, ok := next()
		if !ok {
			break
		}
		println(val)
	}
}
