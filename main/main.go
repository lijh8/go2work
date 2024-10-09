package main

import (
	. "glog2" // https://golang.google.cn/ref/spec#Import_declarations

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	Info()
	Warning()
	Error()
}
