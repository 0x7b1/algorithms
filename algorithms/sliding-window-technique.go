package main

import "fmt"

func slidingWindowBruteforce(s, t string) string {
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s); j++ {
			// if s[i:j] contains all characters of t:
			// update answer
		}
	}

	return ""
}

func main() {
	fmt.Println(slidingWindowBruteforce("ABC", "A"))
}
