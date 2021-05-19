package main

import "fmt"

func countBinarySubstrings(s string) int {
	cur, pre, res := 1, 0 ,0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			if cur < pre {
				res += cur
			} else {
				res += pre
			}

			pre = cur
			cur = 1
		}
	}

	if cur < pre {
		res += cur
	} else {
		res += pre
	}

	return res
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}
