package main

import "fmt"

func lengthOfLastWord(s string) int {
	n, tail := 0, len(s)-1

	for tail >= 0 && s[tail] == ' ' {
		tail--
	}

	for tail >= 0 && s[tail] != ' ' {
		n++
		tail--
	}

	return n
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Hello Tom Tomson"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord(" a "))
}
