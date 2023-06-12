package main

import "fmt"

/**
左右两边子数组和相等
*/

func pivotIndex(nums []int) int {
	sum := 0
	for iii := 0; iii < len(nums); iii++ {
		sum += nums[iii]
	}
	res := -1
	leftSum := 0
	rightSum := sum
	for iii := 0; iii < len(nums); iii++ {
		rightSum -= nums[iii]
		if leftSum == rightSum {
			res = iii
			break
		}
		leftSum += nums[iii]
	}
	return res
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
