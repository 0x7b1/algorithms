package main

import (
	"fmt"
)

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	var digits []int
	var last int

	for x > 0 {
		last = x % 10
		digits = append(digits, last)
		x = x / 10
	}

	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	buff, reverse := x, 0

	for buff > 0 {
		reverse *= 10
		reverse += buff % 10
		buff /= 10
	}

	return x == reverse
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(345543))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1000021))
}
