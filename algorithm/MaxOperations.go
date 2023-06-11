package main

import (
	"fmt"
	"sort"
)

// 双指针
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == k {
			count++
			l++
			r--
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return count
}

// hash
func maxOperations2(nums []int, k int) int {
	var hashmap = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		v, ok := hashmap[n]
		if !ok {
			hashmap[n] = 1
		} else {
			hashmap[n] = v + 1
		}
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		v1, _ := hashmap[temp]
		if v1 <= 0 {
			continue
		}
		diff := k - nums[i]
		v2, ok := hashmap[diff]
		if !ok || v2 <= 0 {
			continue
		}
		if temp == diff && v2 < 2 {
			continue
		}
		count++
		if temp != diff {
			hashmap[temp] = v1 - 1
			hashmap[diff] = v2 - 1
		} else {
			hashmap[diff] = v2 - 2
		}
	}
	return count
}

func maxOperations3(nums []int, k int) int {
	var hashmap = make(map[int]int, len(nums))
	count := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		diff := k - n
		v1, ok := hashmap[diff]
		if ok && v1 > 0 {
			count++
			hashmap[diff] = v1 - 1
		} else {
			v2, ok := hashmap[n]
			if !ok {
				hashmap[n] = 1
			} else {
				hashmap[n] = v2 + 1
			}
		}
	}
	return count
}

func main() {
	fmt.Println(maxOperations2([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations3([]int{3, 1, 3, 4, 3}, 6))
}
