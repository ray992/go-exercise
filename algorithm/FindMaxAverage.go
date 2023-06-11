package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	max := -math.MaxFloat64
	i, j := 0, 0
	var sum = float64(nums[i])
	for {
		if j-i+1 < k {
			j++
			if j == n {
				break
			}
			sum = sum + float64(nums[j])
		} else {
			var temp = sum / float64(j-i+1)
			if temp > max {
				max = temp
			}
			sum = sum - float64(nums[i])
			i++
			if i == n {
				break
			}
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
	fmt.Println(findMaxAverage([]int{-1}, 1))
}
