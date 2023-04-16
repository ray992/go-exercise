package main

import (
	"fmt"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	len := len(piles)
	if len == 1 {
		m := piles[0] / h
		n := piles[0] % h
		if n > 0 {
			return m + 1
		} else {
			return m
		}
	}
	left, right := 1, piles[len-1]
	for left <= right {
		mid := left + (right-left)/2
		sum := 0
		for i := 0; i < len; i++ {
			mod := 0
			if piles[i] > mid && piles[i]%mid > 0 {
				mod = 1
			}
			if piles[i] > mid {
				sum += piles[i]/mid + mod
			} else {
				sum += 1
			}
		}
		if sum > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	//fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	//fmt.Println(minEatingSpeed([]int{30,11,23,4, 20}, 5))
	//fmt.Println(minEatingSpeed([]int{30,11,23,4, 20}, 6))
	//fmt.Println(minEatingSpeed([]int{30,11,23,4, 20}, 12))
	fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
	fmt.Println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
}
