package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	need := make(map[rune]int)

	for _, c := range ransomNote {
		need[c]++
	}

	for _, c := range magazine {
		if _, ok := need[c]; ok {
			need[c]--

			if need[c] == 0 {
				delete(need, c)
			}
		}
	}

	return len(need) == 0
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("a", "aab"))
	fmt.Println(canConstruct("aa", "aab"))
}
