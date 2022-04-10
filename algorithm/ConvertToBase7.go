package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
七进制
*/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negativeNumberStatus := false
	if num < 0 {
		negativeNumberStatus = true
	}
	var res string
	for num != 0 {
		res = res + strconv.Itoa(int(math.Abs(float64(num%7))))
		num = num / 7
	}
	if negativeNumberStatus {
		return "-" + reverseString(res)
	}
	return reverseString(res)
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	//fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(7))
}
