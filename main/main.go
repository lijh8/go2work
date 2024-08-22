package main

type Yield[V any] func(V) bool
type Yield2[K comparable, V any] func(K, V) bool
type Seq[E any] func(Yield[E])
type Seq2[K comparable, V any] func(Yield2[K, V])

// not importing iter, so can't use the named types
func All2[E any](s []E) func(func(E) bool) {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func All2_b[E any](s []E) Seq[E] {
	return func(yield Yield[E]) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	s := []string{"aaa", "bbb", "ccc"}
	for i := range All2(s) {
		println(i)
	}
	for i := range All2_b(s) {
		println(i)
	}

}
