package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i+2] == '#' {
			ss, _ := strconv.Atoi(string(s[i:i+2]))
			res += fmt.Sprintf("%c", 96 + ss)
			i += 2
		} else {
			ss, _ := strconv.Atoi(string(s[i]))
			res += fmt.Sprintf("%c", 96 + ss)
		}
	}

	return res
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}
