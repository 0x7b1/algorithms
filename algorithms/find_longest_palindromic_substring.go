// ---------------------------------------
// Find the Longest Palindromic Substring
// ---------------------------------------
// Example:
// Input: "babad"
// Output: "aba" or "bab"

package main

import "fmt"

// Time: O(n^2), Space: O(1)
func longestPalindromic(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		if len(res) <= len(s1) {
			res = s1
		}

		if len(res) <= len(s2) {
			res = s2
		}
	}

	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}

func main() {
	fmt.Println(longestPalindromic("cbbd"))
	fmt.Println(longestPalindromic("babad"))
}
