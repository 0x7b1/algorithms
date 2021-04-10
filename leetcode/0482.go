package main

import "fmt"

func licenseKeyFormatting(s string, k int) string {
	res := ""
	tmpK := k
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c == '-' {
			continue
		}

		if tmpK == 0 {
			res = "-" + res
			tmpK = k
		}

		tmpK--

		if c >= 'a' && c <= 'z' {
			c = c - 'a' + 'A'
		}

		res = string(c) + res
	}

	return res
}

func main() {
	//fmt.Println(licenseKeyFormatting("AaZz09-", 4))
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}
