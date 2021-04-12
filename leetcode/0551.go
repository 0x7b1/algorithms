package main

import "fmt"

func checkRecord(s string) bool {
	countA, countL := 0, 0
	last := s[0]
	for _, c := range s {
		if c == 'A' {
			countA++
		} else if c == 'L' && last == 'L' {
			countL++
		} else {
			countL = 0
		}

		last = byte(c)
	}

	fmt.Println(countA, countL)

	return countA < 2 && countL <= 3
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
