package main

import "log"

var LOG = log.Println

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
