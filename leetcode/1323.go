package main

import (
	"fmt"
	"math"
)

func maximum69Number(num int) int {
	tmp, d, i6 := num, 0, -1
	for tmp > 0 {
		if tmp%10 == 6 {
			i6 = d
		}

		tmp /= 10
		d++
	}

	if i6 == -1 {
		return num
	}

	return num + int(math.Pow10(i6)*3)
}

func main() {
	fmt.Println(maximum69Number(96969))
}
