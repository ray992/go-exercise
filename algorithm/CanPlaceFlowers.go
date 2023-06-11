package main

import "fmt"

/**
种花问题
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	enableCount := 0
	pre := -1
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			if pre == -1 {
				pre = i
				diff := i - 0
				enableCount += diff / 2
			} else {
				diff := i - pre - 1
				if diff%2 == 0 {
					diff--
				}
				enableCount += diff / 2
				pre = i
			}
		} else {
			if i == len(flowerbed)-1 {
				if pre == -1 {
					enableCount += i/2 + 1
				} else {
					diff := i - pre
					enableCount += diff / 2
				}

			}
		}
	}
	return enableCount >= n
}

func main() {
	var flowerbed = []int{0, 0, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 2))
	// 001 1
	// 0001 1
	// 00001 2
	// 000001 2
	// 0000001 3
	// 00000001 3
	// 000000001 4
}
