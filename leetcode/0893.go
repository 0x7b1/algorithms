package main

import "fmt"

func numSpecialEquivGroups(words []string) int {
	s := make(map[string]bool)

	for _, w := range words {
		odd := make([]int, 26)
		even := make([]int, 26)

		for i, c := range w {
			if i % 2 == 0 {
				even[c-'a']++
			} else {
				odd[c-'a']++
			}
		}

		s[fmt.Sprintf("%v%v", odd, even)] = true
	}

	return len(s)
}

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}
