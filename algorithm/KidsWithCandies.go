package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max = candies[0]
	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	var result = make([]bool, 0, len(candies))
	for i := 0; i < len(candies); i++ {
		if max-candies[i] <= extraCandies {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func main() {
	var candies = []int{4, 2, 1, 1, 2}
	var extraCandies = 1
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
