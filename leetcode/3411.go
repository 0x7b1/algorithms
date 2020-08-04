package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")

	for i := 0; i < len(s)/2; i++ {
		if s[i]!= s[len(s) - i - 1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("anitalavalatina"))
	fmt.Println(isPalindrome("rainforest"))
	fmt.Println(isPalindrome("rfr"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
