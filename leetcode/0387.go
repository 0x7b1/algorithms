package main

import "fmt"

func firstUniqChar(s string) int {
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c]++
	}

	for i, c := range s {
		if chars[c] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
