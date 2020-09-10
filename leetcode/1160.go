package main

import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	mapChars := make(map[rune]int)
	for _, c := range chars {
		mapChars[c]++
	}

	mapCharsBackup := make(map[rune]int)
	for k, v := range mapChars {
		mapCharsBackup[k] = v
	}

	sum := 0
	valid := true
	for _, word := range words {
		valid = true
		for _, c := range word {
			if n := mapChars[c]; n > 0 {
				mapChars[c]--
			} else {
				valid = false
				break
			}
		}

		if valid {
			sum += len(word)
		}

		for k, v := range mapCharsBackup {
			mapChars[k] = v
		}
	}

	return sum
}

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	//fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}
