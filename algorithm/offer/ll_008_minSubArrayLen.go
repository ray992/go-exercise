package main

import "fmt"

// 和大于等于target的最短数组
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	i, j := 0, 0
	minLen := 100001
	sum := 0
	for i <= j && j < length {
		sum += nums[i]
		if sum >= target {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			i++
			j = i
			sum = 0
		} else if sum < target {
			for j = i + 1; j < length; j++ {
				sum += nums[j]
				if sum >= target {
					if j-i+1 < minLen {
						minLen = j - i + 1
					}
					sum = 0
					i++
					j = i
					break
				}
			}
		}
	}
	if minLen == 100001 {
		return 0
	}
	return minLen
}

func minSubArrayLen2(target int, nums []int) int {
	length := len(nums)
	i, j := 0, 0
	minLen := 100001
	sum := 0
	for j < length {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i++
		}
		j++

	}
	if minLen == 100001 {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen2(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen2(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen2(11, []int{1, 2, 3, 4, 5}))
}
