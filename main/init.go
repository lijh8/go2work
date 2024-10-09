package main

import "log"

var LOG = log.Println
var LOGF = log.Printf

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
