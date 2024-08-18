package main

import (
	"strings"
)

func main() {
	haystack := "helloworld"
	needle := "l"
	func(s string) {
		for i, j := 0, 0; i < len(s); i += j + 1 {
			if j = strings.Index(s[i:], needle); j == -1 {
				return
			}
			print(i+j, " ")
		}
	}(haystack)

	println()
	func(s string) {
		for i := len(s); i > 0; {
			if i = strings.LastIndex(s[:i], needle); i != -1 {
				print(i, " ")
			}
		}
	}(haystack)
	println()
}
