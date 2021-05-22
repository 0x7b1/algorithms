package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	countWords := make(map[string]int)

	for _, word := range strings.Split(s1, " ") {
		countWords[word]++
	}

	for _, word := range strings.Split(s2, " ") {
		countWords[word]++
	}

	var uncommon []string
	for word, count := range countWords {
		if count == 1 {
			uncommon = append(uncommon, word)
		}
	}

	return uncommon
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
