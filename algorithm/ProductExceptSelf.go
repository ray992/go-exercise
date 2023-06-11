package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	var l = make([]int, n)
	var r = make([]int, n)
	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = l[i] * r[i]
	}
	return res
}

func main() {
	var nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
