package main

import "fmt"

/** 乘积小于K的字数组 **/

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	var product = 1
	left, right := 0, 0
	ans := 0
	for right < n {
		product *= nums[right]
		for product < k {
			ans++
			right++
			if right >= n {
				break
			}
			product *= nums[right]
		}
		left++
		right = left
		product = 1
	}
	return ans
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	n := len(nums)
	var product = 1
	left, right := 0, 0
	ans := 0
	for right < n {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		ans += right - left + 1
		right++
	}
	return ans
}

func main() {
	fmt.Println(numSubarrayProductLessThanK2([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK2([]int{1, 2, 3}, 0))
}
