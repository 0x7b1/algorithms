package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	lines := 1
	width := 0

	for _, c := range s {
		w := widths[c-'a']
		width += w

		if width > 100 {
			lines++
			width = w
		}
	}

	return []int{lines, width}
}

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}
