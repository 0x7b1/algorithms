package main

import "fmt"

func binaryGap(n int) int {
	last := -1
	res := 0

	for i := 0; i < 32; i++ {
		if ((n >> i) & 1) > 0 {
			if last >= 0 {
				if res < i-last {
					res = i - last
				}
			}

			last = i
		}
	}

	return res
}

func main() {
	fmt.Println(binaryGap(22))
	//fmt.Println(binaryGap(5))
	//fmt.Println(binaryGap(6))
	//fmt.Println(binaryGap(8))
	//fmt.Println(binaryGap(1))
}
