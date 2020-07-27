package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {
	res := make([]int32, len(queries))
	var occurrences int32

	for i, query := range queries {
		for _, s := range strings {
			if query == s {
				occurrences++
			}
		}
		res[i] = occurrences
		occurrences = 0
	}

	return res
}

func main() {
	fmt.Println(matchingStrings([]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}))
}
