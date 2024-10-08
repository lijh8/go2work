package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	m := map[string]int{"aaa": 10, "bbb": 20}
	m["ccc"] = 30
	m["ddd"] = 40
	log.Println(m)

}
