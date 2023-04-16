package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	var window = make(map[byte]int, 8)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		c1 := s[right]
		window[c1] += 1
		right++
		for window[c1] > 1 {
			c2 := s[left]
			window[c2] -= 1
			left++
		}
		ans = int(math.Max(float64(ans), float64(right-left)))
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abccd"))
}
