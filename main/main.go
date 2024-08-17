package main

import (
	"log"
)

func f(s []int) {
	s = append(s, 10, 20, 30)
	log.Println(s)
}

func main() {
	// logfile := "logfile.log"
	// logfd, err := log2.InitLog(logfile)
	// if err == nil {
	// 	defer logfd.Close()
	// }

	s := []int{}
	f(s)
	log.Println(s)
}
