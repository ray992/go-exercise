package main

import (
	"fmt"
	"math"
)

/**
玩筹码
*/
func main() {
	var a []int = []int{1, 2, 3}
	fmt.Println(minCostToMoveChips(a))
}

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for m := 0; m < len(position); m++ {
		if position[m]&1 == 0 {
			even++
		} else {
			odd++
		}
	}
	return int(math.Min(float64(odd), float64(even)))
}
