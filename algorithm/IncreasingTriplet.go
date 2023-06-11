package main

import (
	"fmt"
	"math"
)

/**
递增的三元子序列
*/
func increasingTriplet(nums []int) bool {
	var l = make([]int, len(nums))
	var r = make([]int, len(nums))
	l[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		l[i] = int(math.Min(float64(l[i-1]), float64(nums[i])))
	}
	r[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = int(math.Max(float64(r[i+1]), float64(nums[i+1])))
	}
	for i := 1; i < len(nums); i++ {
		if l[i] < nums[i] && nums[i] < r[i] {
			return true
		}
	}
	return false
}

func increasingTriplet2(nums []int) bool {
	p1 := nums[0]
	p2 := math.MaxUint32
	for i := 1; i < len(nums); i++ {
		if nums[i] > p2 {
			return true
		} else if nums[i] > p1 {
			p2 = nums[i]
		} else {
			p1 = nums[i]
		}
	}
	return false
}

func main() {
	var nums = []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet2(nums))
}
