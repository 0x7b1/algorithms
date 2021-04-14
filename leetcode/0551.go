package main

import "fmt"

func checkRecord(s string) bool {
	countA, countL := 0, 0
	for _, c := range s {
		if c == 'A' {
			countA++
		}

		if c == 'L' {
			countL++
		} else {
			countL = 0
		}

		if countA >= 2 || countL > 2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
