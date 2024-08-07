package main

import (
	"log"
	"log2"
)

func main() {
	logfile := "logfile.log"
	logfd := log2.InitLog(logfile)
	defer logfd.Close()

	log.Println("aaa")
	log.Println("bbb")
}
