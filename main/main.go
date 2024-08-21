package main

type Yield[V any] func(V) bool
type Yield2[K, V any] func(K, V) bool

func sliceDemo[K int, V any](s []V) func(Yield2[K, V]) {
	return func(yield Yield2[K, V]) {
		for i, v := range s {
			if !yield(K(i), v) {
				break
			}
		}
	}
}

func mapDemo[K comparable, V any](s map[K]V) func(Yield2[K, V]) {
	return func(yield Yield2[K, V]) {
		for i, v := range s {
			if !yield(i, v) {
				break
			}
		}
	}
}

func main() {
	s := []string{"aaa", "bbb", "ccc"}
	for i, v := range sliceDemo(s) {
		println(i, v)
	}

	m := map[string]int{"aaa": 10, "bbb": 20, "ccc": 30}
	for i, v := range mapDemo(m) {
		println(i, v)
	}
}
