package main

import "fmt"

/*
func main() {
	// synchronize two goroutines with two channels
	// let the code in the first goroutine execute first
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
}
*/

func main() {
	// synchronize two goroutines with two channels
	// let the code in the second goroutine execute first
	const MAX = 6
	cA := make(chan any)
	cB := make(chan any)
	done := make(chan any)

	go func() {
		for i := 0; i < MAX; i++ {
			<-cA
			fmt.Println("a", i)
			cB <- nil
		}
		done <- nil
	}()

	go func() {
		for i := 0; i < MAX; i++ {
			fmt.Println("b", i)
			cA <- nil
			<-cB
		}
		done <- nil
	}()

	<-done
	<-done
}
