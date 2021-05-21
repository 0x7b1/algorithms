package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	freqMap, start := make(map[string]int), -1
	for i, c := range paragraph {
		if c == ' ' || c == '!' || c == '?' || c == '\'' || c == ',' || c == ';' || c == '.' {
			if start > -1 {
				word := strings.ToLower(paragraph[start:i])
				freqMap[word]++
			}
			start = -1
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		word := strings.ToLower(paragraph[start:])
		freqMap[word]++
	}
	for _, bannedWord := range banned {
		delete(freqMap, bannedWord)
	}
	mostFreqWord, mostFreqCount := "", 0
	for word, freq := range freqMap {
		if freq > mostFreqCount {
			mostFreqWord = word
			mostFreqCount = freq
		}
	}
	return mostFreqWord
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"}))
}
