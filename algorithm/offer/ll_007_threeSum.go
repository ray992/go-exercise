package main

import "sort"

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ll := len(nums)
	result := make([][]int, 0)
	for ii := 0; ii < ll; ii++ {
		if ii > 0 && nums[ii] == nums[ii-1] {
			continue
		}
		f1 := nums[ii]
		third := ll - 1
		for jj := ii + 1; jj < ll; jj++ {
			if jj > ii+1 && nums[jj] == nums[jj-1] {
				continue
			}
			f2 := nums[jj]
			for jj < third && f1+f2+nums[third] > 0 {
				third--
			}
			if jj == third {
				break
			}
			if f1+f2+nums[third] == 0 {
				result = append(result, []int{nums[ii], nums[jj], nums[third]})
			}
		}
	}
	return result
}

func main() {

}
