package main

import "fmt"

func shuffle(nums []int, n int) []int {
	tmp := make([]int, len(nums))
	for i := 0; i < n; i++ {
		tmp[2*i], tmp[2*i+1] = nums[i], nums[n+i]
	}

	return tmp
}

func main() {
	a := []int{2,5,1,3,4,7}
	fmt.Println(shuffle(a, 3))
}
