package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, last := 0, len(nums)-1

	for i <= last {
		if nums[i] == val {
			nums[i], nums[last] = nums[last], nums[i]
			last--
		} else {
			i++
		}
	}

	return last + 1
}

// 3 2 2 3
// 0 1 2 3

func main() {
	a1 := []int{3, 2, 2, 3}
	fmt.Println(a1)
	la1 := removeElement(a1, 3)
	fmt.Println(la1, a1)
}
