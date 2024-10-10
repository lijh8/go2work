package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	m := 0
	for i, j := 0, 0; j < len(s); j++ {
		if k := strings.IndexByte(s[i:j], s[j]); k != -1 {
			i = i + k + 1
		}
		m = max(m, j-i+1)
	}
	return m
}

func main() {
	LOG(lengthOfLongestSubstring("abcabcbb")) // 3
	LOG(lengthOfLongestSubstring("bbbbb"))    // 1
	LOG(lengthOfLongestSubstring("pwwkew"))   // 3
}
