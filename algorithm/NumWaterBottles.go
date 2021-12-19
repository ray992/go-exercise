package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles <  numExchange {
		return numBottles
	}
	sum := numBottles
	changeCount := numBottles
	mod := 0
	for (changeCount + mod) >= numExchange{
		changeCount = numBottles / numExchange
		mod = numBottles % numExchange
		sum = sum + changeCount
		numBottles = changeCount + mod
	}
	return sum
}

func main() {
	fmt.Println(numWaterBottles(9,3))
}
