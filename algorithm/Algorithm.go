package main

import "fmt"

func runningSum(nums []int) []int {
	n := len(nums)
	var res []int
	temp := 0
	for i := 0; i < n; i++ {
		temp += nums[i]
		res = append(res, temp)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
	nums2 := []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(nums2))
}
