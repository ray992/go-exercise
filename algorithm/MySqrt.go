package main

import "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		mul := mid * mid
		if mul == x {
			return mid
		} else if mul > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(15))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(100))
}
