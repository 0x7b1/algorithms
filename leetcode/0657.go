package main

import "fmt"

func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, c := range moves {
		if c == 'L' {
			x--
		} else if c == 'R' {
			x++
		} else if c == 'U' {
			y++
		} else if c == 'D' {
			y--
		}
	}

	return x == 0 && y == 0
}

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}
