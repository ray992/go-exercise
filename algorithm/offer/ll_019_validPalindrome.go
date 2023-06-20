package main

import "fmt"

/**
最多删除一个字符得到回文
*/

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			newLeft := left + 1
			newRight := right - 1
			b1 := delOneByteValidPalindrome(s, newLeft, right) // 核心就是嵌套处理
			b2 := delOneByteValidPalindrome(s, left, newRight)
			return b1 || b2
		}
	}
	return true
}

func delOneByteValidPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	//fmt.Println(validPalindrome("aba"))
	//fmt.Println(validPalindrome("abc"))
	//fmt.Println(validPalindrome("abca"))
	//fmt.Println(validPalindrome("eedede"))
	fmt.Println(validPalindrome("pidbliassaqozokmtgahluruufwbjdtayuhbxwoicviygilgzduudzgligyviciowxbhuyatdjbwfuurulhagtmkozoqassailbdip"))
}
