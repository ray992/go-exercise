package main

import "fmt"

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var diff = 0
	var res = 0
	for i < j {
		if height[j] >= height[i] {
			diff = height[i]
			i++
		} else {
			diff = height[j]
			j--
		}
		temp := diff * (j - i + 1)
		if temp > res {
			res = temp
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
