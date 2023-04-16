package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		p := left + (right-left)/2
		if nums[p] == target {
			return p
		} else if nums[p] > target {
			right = p - 1
		} else {
			left = p + 1
		}
	}
	//fmt.Println(left, right)
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 10))
	fmt.Println(searchInsert([]int{1, 3}, 2))
}
