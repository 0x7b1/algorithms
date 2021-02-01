package main

import "fmt"

func stringContains(a, b string) bool {
	counter := make(map[rune]int)
	for _, c := range a {
		counter[c]++
	}

	for _, c := range b {
		if _, ok := counter[c]; !ok {
			return false
		}

		counter[c]--

		if counter[c] < 0 {
			return false
		}
	}

	return true
}

// Time: O(n^2)
func minimumWindowSubstringBruteforce(s, t string) string {
	var minSubstr string = s
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			substr := s[i : j+1]
			if stringContains(substr, t) {
				if len(substr) < len(minSubstr) {
					minSubstr = substr
				}
			}
		}
	}

	return minSubstr
}

// Time: O(n)
func minimumWindowSubstringSliding(s, t string) string {
	minLenght := 10000
	start := 0
	var left, right int
	windowHave := make(map[byte]int)
	windowNeed := make(map[byte]int)
	match := 0

	for _, c := range t {
		windowNeed[byte(c)]++
	}

	for right < len(s) {
		c1 := s[right]

		if windowNeed[c1] > 0 {
			windowHave[byte(c1)]++

			if windowHave[c1] == windowNeed[c1] {
				match++
			}
		}

		right++

		for match == len(windowNeed) {
			if right-left < minLenght {
				start = left
				minLenght = right - left
			}

			c2 := s[left]

			if windowNeed[c2] > 0 {
				windowHave[c2]--

				if windowHave[c2] < windowNeed[c2] {
					match--
				}
			}
			left++
		}
	}

	return s[start : start+minLenght]
}

func findAllAnagramsInString(s, t string) []int {
	var left, right int
	windowHave := make(map[byte]int)
	windowNeed := make(map[byte]int)
	match := 0
	res := []int{}

	for _, c := range t {
		windowNeed[byte(c)]++
	}

	for right < len(s) {
		c1 := s[right]

		if windowNeed[c1] > 0 {
			windowHave[byte(c1)]++

			if windowHave[c1] == windowNeed[c1] {
				match++
			}
		}

		right++

		for match == len(windowNeed) {
			if right-left == len(t) {
				res = append(res, left)
			}

			c2 := s[left]

			if windowNeed[c2] > 0 {
				windowHave[c2]--

				if windowHave[c2] < windowNeed[c2] {
					match--
				}
			}
			left++
		}
	}

	return res
}

func lengthOfLongestSubstringSliding(s string) int {
	window := make(map[byte]int)
	var left, right int
	var res int

	for right < len(s) {
		c1 := s[right]
		window[c1]++
		right++

		if window[c1] > 1 {
			c2 := s[left]
			window[c2]--
			left++
		}

		if res < right-left {
			res = right - left
		}
	}

	return res
}

func main() {
	fmt.Println(minimumWindowSubstringBruteforce("ABCB", "BB"))
	fmt.Println(minimumWindowSubstringSliding("ABCB", "BB"))

	fmt.Println(minimumWindowSubstringBruteforce("EBBANCF", "ABC"))
	fmt.Println(minimumWindowSubstringSliding("EBBANCF", "ABC"))

	fmt.Println(findAllAnagramsInString("cbaebabacd", "abc"))

	fmt.Println(lengthOfLongestSubstringSliding("abcabcbb"))
}
