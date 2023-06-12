package main

import "fmt"

// 和为K的子数组
func subarraySum(nums []int, k int) int {
	n := len(nums)
	ans := 0
	for start := 0; start < n; start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

func subarraySum2(nums []int, k int) int {
	n := len(nums)
	ans := 0
	numMap := make(map[int]int)
	numMap[0] = 1
	preSum := 0
	for start := 0; start < n; start++ {
		preSum += nums[start]
		v, ok := numMap[preSum-k]
		if ok {
			ans += v
		}
		v1, ok := numMap[preSum]
		if ok {
			numMap[preSum] = v1 + 1
		} else {
			numMap[preSum] = 1
		}
	}
	return ans
}

func main() {
	fmt.Println(subarraySum2([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum2([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum2([]int{-1, -1, 1}, 0))
	fmt.Println(subarraySum2([]int{1, -1, 0}, 0))
}
