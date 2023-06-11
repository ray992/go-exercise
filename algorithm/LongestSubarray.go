package main

import "fmt"

func longestSubarray(nums []int) int {
	n := len(nums)
	ans := 0
	p0, p1 := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			p1 = p0
			p0 = 0
		} else {
			p0++
			p1++
		}
		if p1 > ans {
			ans = p1
		}
	}
	if ans == n {
		ans--
	}
	return ans
}

func main() {
	//fmt.Println(longestSubarray([]int{1,1,0,1}))
	//fmt.Println(longestSubarray([]int{0,1,1,1,0,1,1,0,1}))
	//fmt.Println(longestSubarray([]int{1,1,1}))
	fmt.Println(longestSubarray([]int{1, 1, 0, 0, 1, 1, 1, 0, 1}))
}
