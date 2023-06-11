package main

import "fmt"

func largestAltitude(gain []int) int {
	n := len(gain)
	latitude := make([]int, n+1)
	latitude[0] = 0
	pre := 0
	for i := 0; i < n; i++ {
		pre = pre + gain[i]
		latitude[i+1] = pre
	}
	max := latitude[0]
	for i := 1; i < len(latitude); i++ {
		if max < latitude[i] {
			max = latitude[i]
		}
	}
	return max
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
