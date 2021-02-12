package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if len(stack) > 0 && stack[len(stack) - 1] == c {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("]"))
}
