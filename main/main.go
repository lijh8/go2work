package main

import (
	"log2"
	"tuple2"
)

func main() {
	filename := "logfile.log"
	logfd := log2.InitLog(filename)
	defer logfd.Close()

	s1 := []any{1, "hello", 3.14, true}
	s2 := []any{1, "hello", 3.14, true}
	b, err := tuple2.Cmp(s1, s2)
	println(b, err == nil)

}
