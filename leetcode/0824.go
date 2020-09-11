package main

import (
	"fmt"
	"strings"
	"unicode"
)

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	var phrase strings.Builder

	for i, word := range words {
		first := unicode.ToLower(rune(word[0]))

		if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' {
			word += "ma"
		} else {
			word = word[1:] + string(word[0]) + "ma"
		}

		phrase.WriteString(word + strings.Repeat("a", i+1) + " ")
	}

	res := phrase.String()
	res = res[:len(res) - 1]

	return res
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}
