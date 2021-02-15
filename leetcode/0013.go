package main

import "fmt"

func romanToInt(s string) int {
	var res int

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			if res >= 5 {
				res -= 1
			} else {
				res += 1
			}
		case 'V':
			res += 5
		case 'X':
			if res >= 50 {
				res -= 10
			} else {
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			if res >= 500 {
				res -= 100
			} else {
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}

	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
}
