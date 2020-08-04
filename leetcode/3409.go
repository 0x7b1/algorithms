package main

import "fmt"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	valid := true

	first := word[0]
	rest := word[1:]

	// check if the rest is all upper case
	for _, w := range rest {
		if w > 96 { // the first lower case
			valid = false
			break
		}
	}

	if first < 96 && valid {
		return valid
	}

	if first > 96 && rest[0] < 96 {
		return false
	}

	if !valid {
		// check if the rest is all lower case
		for _, w := range rest {
			if w > 96 {
				valid = true
				// break
			} else {
				valid = false
				break
			}
		}
	}

	return valid
}


func main() {
	fmt.Println(detectCapitalUse("Pikachu"))
}
