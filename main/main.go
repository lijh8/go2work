package main

import "fmt"

func main() {
	ab()
	ba()
}

// show "a","b" alternatively with two goroutines
func ab() {
	const MAX = 6
	cA := make(chan any)
	cB := make(chan any)
	done := make(chan any)

	go func() {
		for i := 0; i < MAX; i++ {
			fmt.Println("a", i)
			cA <- nil
			<-cB
		}
		done <- nil
	}()

	go func() {
		for i := 0; i < MAX; i++ {
			<-cA
			fmt.Println("b", i)
			cB <- nil
		}
		done <- nil
	}()

	<-done
	<-done
	fmt.Println("---")
}

// show "b","a" alternatively with two goroutines
func ba() {
	const MAX = 6
	cA := make(chan any)
	cB := make(chan any)
	done := make(chan any)

	go func() {
		for i := 0; i < MAX; i++ {
			fmt.Println("b", i)
			cA <- nil
			<-cB
		}
		done <- nil
	}()

	go func() {
		for i := 0; i < MAX; i++ {
			<-cA
			fmt.Println("a", i)
			cB <- nil
		}
		done <- nil
	}()

	<-done
	<-done
	fmt.Println("---")
}
