package main

import (
	"fmt"
)

// 只出线一次的数字

func singleNumber(nums []int) int {
	var numMap = make(map[int]int)
	var result int
	for i := 0; i < len(nums); i++ {
		v, ok := numMap[nums[i]]
		if ok {
			numMap[nums[i]] = v + 1
		} else {
			numMap[nums[i]] = 1
		}
	}
	for k, v := range numMap {
		if v == 1 {
			result = k
			break
		}
	}
	return result
}

func singleNumber2(nums []int) int {
	cnt := [32]int32{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(cnt); j++ {
			if ((nums[i] >> j) & 1) == 1 {
				cnt[j]++
			}
		}
	}
	ans := int32(0)
	for i := 0; i < 32; i++ {
		if ((cnt[i] % 3) & 1) == 1 {
			ans += 1 << i
		}
	}
	return int(ans)
}

func main() {
	//fmt.Println(singleNumber([]int{2,2,3,2}))
	//fmt.Println(singleNumber([]int{0,1,0,1,0,1,100}))
	//fmt.Println(singleNumber2([]int{30000,500,100,30000,100,30000,100}))
	fmt.Println(singleNumber2([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
