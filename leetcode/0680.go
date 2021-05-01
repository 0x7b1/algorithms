package main

import "fmt"

func isPalindrome(str string, s, t int) bool {
	for s <= t {
		if str[s] == str[t] {
			s++
			t--
		} else {
			return false
		}
	}

	return true
}

func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return isPalindrome(s, l, r-1) || isPalindrome(s, l+1, r)
		}
	}

	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
}
