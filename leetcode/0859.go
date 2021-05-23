package main

import "fmt"

func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	setA := make(map[rune]bool)
	setB := make(map[rune]bool)

	for _, c := range a {
		setA[c] = true
	}

	for _, c := range b {
		setB[c] = true
	}

	if len(setA) != len(setB) {
		return false
	}

	if a == b {
		return len(a) - len(setA) >= 1
	} else {
		indices := []int{}
		counter := 0

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter++
				indices = append(indices, i)
			}

			if counter > 2 {
				return false
			}
		}

		if counter != 2 {
			return false
		}

		return a[indices[0]] == b[indices[1]] && a[indices[1]] == b[indices[0]]
	}
}

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aa", "ax"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
}
