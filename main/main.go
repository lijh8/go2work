package main

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	incr := increment()
	println(incr())
	println(incr())
	println(incr())
}
