package main

import "fmt"

/**
	回文子字符串个数
**/

func countSubstrings(s string) int {
	n := len(s)
	left, right := 0, 0
	count := 0
	for left < n {
		if hw(s, left, right) {
			count++
		}

		right++
		if right == n {
			left++
			right = left
		}
		//fmt.Println(left, right, count)
	}
	return count
}

/** DP实现 **/
func countSubstrings2(s string) int {
	n := len(s)
	count := 0
	var dp [1001][1001]bool
	for ii := n - 1; ii >= 0; ii-- {
		for jj := ii; jj < n; jj++ {
			if s[ii] == s[jj] && (jj-ii <= 1 || dp[ii+1][jj-1]) {
				count++
				dp[ii][jj] = true
			}
		}
	}
	return count
}

func hw(s string, left, right int) bool {
	for i := left; i <= right; i++ {
		if s[i] != s[right-(i-left)] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("a"))
	fmt.Println(countSubstrings("adbaabda"))
	fmt.Println(countSubstrings("aba"))
}
