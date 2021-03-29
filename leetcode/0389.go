package main

import "fmt"

func findTheDifferenceMap(s string, t string) byte {
	chars := make(map[byte]int)

	for _, c := range t {
		chars[byte(c)]++
	}

	for _, c := range s {
		if _, ok := chars[byte(c)]; ok {
			chars[byte(c)]--

			if chars[byte(c)] == 0 {
				delete(chars, byte(c))
			}
		}

	}

	var val byte
	for c := range chars {
		val = c
		break
	}

	return val
}

func findTheDifference(s string, t string) byte {
	n, ch := len(t), t[len(t) - 1]

	for i := 0; i < n-1; i++ {
		ch ^= s[i]
		ch ^= t[i]
	}

	return ch
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
	fmt.Println(findTheDifference("", "y"))
	fmt.Println(findTheDifference("xyz", "xpyz"))
	fmt.Println(findTheDifference("a", "aa"))
	fmt.Println(findTheDifference("ae", "aea"))
}
