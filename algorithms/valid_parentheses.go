package main

import "fmt"

func isValid(str string) bool {
	var left_stack []rune

	for _, c := range str {
		if c == '[' || c == '(' || c == '{' {
			left_stack = append(left_stack, c)
		} else {
			if len(left_stack) > 0 && leftOf(c) == left_stack[len(left_stack)-1] {
				left_stack = left_stack[:len(left_stack)-1]
			} else {
				return false
			}
		}

	}

	return len(left_stack) == 0
}

func leftOf(c rune) rune {
	if c == ']' {
		return '['
	} else if c == ')' {
		return '('
	}

	return '{'
}

func main() {
	fmt.Println(isValid("[[]]"))
	fmt.Println(isValid("{()}"))
	fmt.Println(isValid("{}[]()"))
	fmt.Println(isValid("[(])"))
}
