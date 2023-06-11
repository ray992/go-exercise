package main

import "fmt"

// 判断子序列
func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		exist := false
		for j < len(t) {
			if b == t[j] {
				exist = true
				j++
				break
			}
			j++
		}
		if !exist {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
