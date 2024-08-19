package main

import (
	"log"
	"log2"
)

func main() {
	logfile := "logfile.log"
	logfd, err := log2.InitLog(logfile)
	if err == nil {
		defer logfd.Close()
	}

	log.Println("aaa")
	log.Println("bbb")

}
