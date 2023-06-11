package main

import "fmt"

/**
字符串的最大共因子
*/
func gcdOfStrings(str1 string, str2 string) string {
	var s1 = str1 + str2
	var s2 = str2 + str1
	if s1 != s2 {
		return ""
	}
	return s1[0:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	fmt.Println(gcdOfStrings("abc", "abcabc"))
	fmt.Println(gcdOfStrings("abcabc", "abcabc"))
}
