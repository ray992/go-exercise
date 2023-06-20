package main

import "fmt"

/**
最长不重复的子串
*/
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	left := 0
	maxLength := 0
	byteMap := make(map[byte]bool)
	for right := 0; right < n; right++ {
		b := s[right]
		for {
			_, ok := byteMap[b]
			if ok {
				delete(byteMap, s[left])
				left++
			} else {
				break
			}
		}
		byteMap[b] = true
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
