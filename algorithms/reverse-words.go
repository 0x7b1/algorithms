package main

import "fmt"

func reverseWords(sentence []byte) {
	var start, end int

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' || i == len(sentence)-1 {
			end = i - 1

			if i == len(sentence)-1 {
				end = i
			}

			for start < end {
				sentence[start], sentence[end] = sentence[end], sentence[start]
				start++
				end--
			}

			start = i + 1
		}
	}
}

func main() {
	sentence := []byte{'I', ' ', 'l', 'o', 'v', 'e'}
	reverseWords(sentence)
	fmt.Println(string(sentence))
}
