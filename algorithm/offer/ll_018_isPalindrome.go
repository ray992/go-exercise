package main

import "strings"

/**
有效的回文
*/
func isPalindrome2(s string) bool {
	ss := "abcdefghijklmnopqrstuvwxyz0123456789"
	vm := make(map[byte]bool)
	for i := 0; i < len(ss); i++ {
		vm[ss[i]] = true
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		b1 := s[left]
		b2 := s[right]
		_, ok := vm[b1]
		if !ok {
			left++
			continue
		}
		_, ok = vm[b2]
		if !ok {
			right--
			continue
		}
		if b1 != b2 {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

// 双指针法
func isPalindrome3(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		b1 := s[left]
		b2 := s[right]
		if !isValid(b1) {
			left++
			continue
		}
		if !isValid(b2) {
			right--
			continue
		}
		if b1 != b2 {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {

}
