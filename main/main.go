package main

import (
	"log"
	"log2"
	"os"
	"path/filepath"
)

func main() {
	logpath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	logpath = filepath.Dir(logpath)
	logpath = filepath.Join(logpath, "logfile.log")
	logfd := log2.InitLog(logpath)
	defer logfd.Close()

	log.Println("aaa")
	log.Println("bbb")
}
