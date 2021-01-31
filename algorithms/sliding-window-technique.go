package main

import "fmt"

func minSlidingWindowBruteforce(s, t string) string {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// if s[i:j] contains all characters of t:
			// update answer
		}
	}

	return ""
}

func minSlidingWindowIncremental(s, t string) string {
	var start, left, right int
	minLength := 1000
	var match int
	window := make(map[byte]int)
	needs := make(map[byte]int)

	for _, c := range t {
		needs[byte(c)]++
	}

	for right < len(s) {
		c1 := s[right]
		if count, ok := needs[c1]; ok && count > 0 {
			window[c1]++
			if window[c1] == needs[c1] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left < minLength {
				start = left
				minLength = right - left
			}

			c2 := s[left]
			if count, ok := needs[c2]; ok && count > 0 {
				window[c2]--
				if window[c2] < needs[c2] {
					match--
				}
			}
			left++
		}
	}

	return s[start:minLength]
}

func main() {
	fmt.Println(minSlidingWindowBruteforce("ABC", "A"))
}
