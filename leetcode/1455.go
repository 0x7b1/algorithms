package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.Index(word, searchWord) == 0 {
			return i + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
}
