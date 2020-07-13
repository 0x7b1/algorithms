package main

import "fmt"

func removeOuterParentheses(S string) string {
	deep, result, ctr := 1, "", 0
	for _, c := range S {
		if c == '(' && ctr >= deep {
			result += string(c)
		}
		if c == ')' && ctr > deep {
			result += string(c)
		}

		if c == '(' {
			ctr++
		} else {
			ctr--
		}
	}

	return result
}

func main() {
	fmt.Println(removeOuterParentheses("()()"))
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
