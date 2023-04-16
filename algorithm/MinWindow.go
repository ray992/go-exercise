package main

import (
	"fmt"
	"math"
)

/**

 */
func minWindow(s string, t string) string {
	var need [256]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	l, r := 0, 0
	size := math.MaxInt32
	count := len(t)
	start := 0
	for r < len(s) {
		if need[s[r]] > 0 {
			count--
		}
		need[s[r]]--
		if count == 0 {
			for l < r && need[s[l]] < 0 {
				need[s[l]]--
				l++
			}
			if r-l+1 < size {
				start = l
				size = r - l + 1
			}
			need[s[l]]++
			l++
			count++
		}
		r++
	}
	if size == math.MaxInt32 {
		return ""
	}
	return s[(start):(start + size)]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	//fmt.Println(minWindow("aa", "aa"))
	//fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
