package main

import "fmt"

// This needs debugging

func isMatchBruteforce(text, pattern string) bool {
	if len(pattern) == 0 {
		return len(text) > 0
	}

	first := len(text) > 0 && (pattern[0] == text[0] || pattern[0] == '.')

	if len(pattern) >= 2 && pattern[1] == '*' {
		return isMatchBruteforce(text, pattern[2:]) || first && isMatchBruteforce(text[1:], pattern)
	}

	return first && isMatchBruteforce(text[1:], pattern[1:])
}

func isMatchDP(text, pattern string) bool {
	memo := make([][]bool, len(text) * 2)
	for i := range memo {
		memo[i] = make([]bool, len(text) * 2)
	}
	var dp func(i, j int) bool

	dp = func(i, j int) bool {
		if memo[i][j] {
			return memo[i][j]
		}

		if j == len(pattern) {
			return i == len(text)
		}

		first := len(text) > 0 && (pattern[0] == text[0] || pattern[0] == '.')

		var ans bool
		if j <= len(pattern)-2 && pattern[j+1] == '*' {
			ans = dp(i, j+2) || first && dp(i+1, j)
		} else {
			ans = first && dp(i+1, j+1)
		}

		memo[i][j] = ans
		return ans
	}

	return dp(0, 0)
}

func main() {
	fmt.Println(isMatchBruteforce("aab", "a*b"))
	fmt.Println(isMatchDP("aab", "a*b"))
}
