package main

import "fmt"

func removeDuplicates(S string) string {
	result := ""
	for _, c := range S {
		if len(result) > 0 && c == rune(result[len(result) - 1]) {
			result = result[:len(result) - 1]
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
