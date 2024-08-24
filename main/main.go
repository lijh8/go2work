package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	t := []int{10, 20, 30, 40, 50}
	copy(a[:], t)
	// copy(a, t[:]) //no
	copy(t, a[:]) //no
	fmt.Println(t, a)
}
