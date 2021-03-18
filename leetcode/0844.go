package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	var stackS, stackT []rune

	for _, c := range S {
		if len(stackS) > 0 && c == '#' {
			stackS = stackS[:len(stackS)-1]
		}
		if c != '#' {
			stackS = append(stackS, c)
		}
	}

	for _, c := range T {
		if len(stackT) > 0 && c == '#' {
			stackT = stackT[:len(stackT)-1]
		}

		if c != '#' {
			stackT = append(stackT, c)
		}
	}

	return string(stackS) == string(stackT)
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
}
