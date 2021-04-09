package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	ss := s + s
	ss = ss[1:len(ss) - 1]

	return strings.Contains(ss, s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
