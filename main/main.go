package main

import (
	"log"
	"log2"
)

func main() {
	filename := "logfile.log"
	logfd := log2.InitLog(filename)
	defer logfd.Close()

	log.Println("aaa")
	log.Println("bbb")
}
