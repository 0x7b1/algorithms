package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	patternByte := []byte(pattern)
	if pattern == "" || len(patternByte) != len(words) {
		return false
	}

	pMap := map[byte]string{}
	sMap := map[string]byte{}

	for i, b := range patternByte{
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[words[i]]; !ok {
				pMap[b] = words[i]
				sMap[words[i]] = b
			} else {
				if sMap[words[i]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != words[i] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
