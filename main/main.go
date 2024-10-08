package main

import (
	"log"
	"reflect"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := "abc"
	var s2 string = "abcde"
	log.Println(reflect.TypeOf(s))
	log.Println(reflect.TypeOf(s2))
}
