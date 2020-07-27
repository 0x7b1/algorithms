package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) (rev string) {
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

func reverseWords(s string) (rev string) {
	words := strings.Split(s, " ")
	for _, word := range words {
		rev += reverseWord(word) + " "
	}

	return rev[:len(rev) - 1]
}

func main() {
	fmt.Println(reverseWords("Hello World"))
}
