package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	res := ""
	c := 0
	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || c == 1 {
		if i >= 0 {
			c += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			c += int(b[j] - '0')
			j--
		}

		res = string(rune(c%2+'0')) + res

		c /= 2
	}

	return res
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "0"))
}
