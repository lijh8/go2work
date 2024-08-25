package main

import (
	"os"
	"os/signal"
	"syscall"
)

// $ kill -HUP <pid>
// $ kill -HUP <pid>

func handlerHUP(sig os.Signal) {
	println("handlerHUP: ", sig == syscall.SIGHUP)
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	done := make(chan struct{})

	go func() {
		for sig := range ch {
			handlerHUP(sig)
		}
		done <- struct{}{}
	}()

	//...

	println("awaiting SIGHUP")

	<-done
	println("exiting")
}

/*
$ go build && ./main
awaiting SIGHUP
handlerHUP:  true
handlerHUP:  true
^C
$
*/
