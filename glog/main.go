package main

import "github.com/golang/glog"

func main() {
	InitLog() // in same package
	defer glog.Flush()

	for {
		glog.Info("INFO!")
		glog.Warning("WARNING!")
		glog.Error("ERROR!")
		// glog.Fatal("FATAL!") // this will exit the program
	}
}
