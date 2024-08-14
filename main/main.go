package main

// var c = make(chan int) // ok
var c = make(chan int, 10) // no, buffered
var a string

var done chan any = make(chan any)

func f() {
	a = "hello, world"
	println("len:", len(c), ",", "cap:", cap(c)) // values in c is sent in main()
	<-c
	done <- nil
}

func main() {
	go f()
	<-done
	c <- 0 // try comment out this line
	print(a)

}
