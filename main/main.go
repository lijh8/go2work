package main

import (
	"fmt"
)

func main() {
	// logfile := "logfile.log"
	// logfd, err := log2.InitLog(logfile)
	// if err == nil {
	// 	defer logfd.Close()
	// }

	// log.Println("aaa")
	// log.Println("bbb")

	cnt := 0
	go func() {
		for {
			cnt++
			fmt.Print(cnt, ",")
		}
	}()
	select {} // Block forever

}
