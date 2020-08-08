package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	res := []string{}

	if len(words) < 3 {
		return res
	}

	for i := 0; i < len(words) - 2; i++ {
		if words[i] == first && words[i + 1] == second {
			res = append(res, words[i + 2])
		}
	}

	return res
}

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}
