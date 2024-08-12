package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	done := make(chan struct{})

	//sending
	var sending = func() {
		for i := 0; i <= 5; i++ {
			select {
			case ch <- i:
				fmt.Printf("sending:%d, ", i)
			default:
				// block on channel is full, unless there is a receiver
				fmt.Printf("blocking, len:%d, cap:%d, ", len(ch), cap(ch))
			}
		}

		close(ch)
		done <- struct{}{}
	}

	//receiving
	var receiving = func() {
		// turn of this range loop to see the block when channel is full
		for i := range ch {
			fmt.Printf("receiving, ch:%d, len:%d, cap:%d, ", i, len(ch), cap(ch))
		}

		done <- struct{}{}
	}

	go sending()
	go receiving()

	<-done // one for sending
	<-done // one for receiving
}
