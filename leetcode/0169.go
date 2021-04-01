package main

import "fmt"

func majorityElement2(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	var maxCount, maxNum int
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			maxNum = num
		}
	}

	return maxNum
}

func majorityElement(nums []int) int {
	res, count := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}

	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
