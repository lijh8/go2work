package main

import (
	"log"
	"log2"
)

func main() {
	filename := "logfile.log"
	logfd, err := log2.InitLog(filename)
	if err == nil {
		defer logfd.Close()
	}

	log.Println("aaa")
	log.Println("bbb")
}
