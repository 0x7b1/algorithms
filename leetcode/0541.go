package main

import "fmt"

func reverseStr(s string, k int) string {
	a := []byte(s)
	for start := 0; start < len(a); start +=2*k {
		i := start
		var j = start + k - 1
		if len(a) - 1 < j {
			j = len(s) - 1
		}

		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	return string(a)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}
