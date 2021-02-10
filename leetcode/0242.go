package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make(map[rune]int)

	for _, c := range s {
		chars[c]++
	}

	for _, c := range t {
		if _, ok := chars[c]; !ok {
			return false
		}

		chars[c]--
		if chars[c] == 0 {
			delete(chars, c)
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("aacc", "ccac"))
}
