package main

var c = make(chan int, 3)
var a string
var done = make(chan any)

func f() {
	a = "hello, world"
	println("aaa", len(c), cap(c))
	v := <-c

	println("bbb", len(c), cap(c), v)

	done <- nil
}

func main() {
	go f()
	c <- 10
	c <- 20

	println("ccc", a)
	<-done
}
