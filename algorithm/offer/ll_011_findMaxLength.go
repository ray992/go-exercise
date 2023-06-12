package main

import "fmt"

/**
0和1个数相同的子数组 ， 前缀和 O（n）
*/
func findMaxLength(nums []int) int {
	n := len(nums)
	maxLength := 0
	counter := 0
	counterMap := make(map[int]int)
	counterMap[counter] = -1
	for ii := 0; ii < n; ii++ {
		curNum := nums[ii]
		if curNum == 1 {
			counter++
		} else {
			counter--
		}
		v, ok := counterMap[counter]
		if ok {
			if ii-v > maxLength {
				maxLength = ii - v
			}
		} else {
			counterMap[counter] = ii
		}
	}
	return maxLength
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}
