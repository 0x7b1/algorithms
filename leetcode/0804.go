package main

import "fmt"

var codes = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	var morse_word string
	morse_codes := make(map[string]bool)

	for _, word := range words {
		for _, c := range word {
			morse_word += codes[c - rune('a')]
		}

		morse_codes[morse_word] = true
		morse_word = ""
	}


	return len(morse_codes)
}

// "a" -> ".-"
// "b" -> "-..."

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
