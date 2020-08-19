package main

import "fmt"

func diStringMatch(S string) []int {
	increase, decrease := 0, len(S)
	var out []int

	for _, c := range S {
		if c == 'D' {
			out = append(out, decrease)
			decrease--
		} else if c == 'I' {
			out = append(out, increase)
			increase++
		}
	}

	out = append(out, increase)

	return out
}

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
}
