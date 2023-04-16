package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left := 0
	right := n - 1
	for left <= right {
		middle := (right-left)/2 + left
		var cur byte = letters[middle]
		if cur > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if right+1 >= n {
		return letters[0]
	}
	return letters[right+1]
}

func main() {
	var l = []byte{'a', 'b', 'c'}
	fmt.Println(nextGreatestLetter(l, 'd'))
}
