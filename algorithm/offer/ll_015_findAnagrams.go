package main

import "fmt"

/** 字符串中的所有变位词 **/
func findAnagrams(s string, p string) []int {
	var res []int
	var arr2 = [26]int{}
	for i := 0; i < len(p); i++ {
		arr2[p[i]-'a']++
	}
	left, right := 0, len(p)-1
	for right < len(s) {
		var arr1 = [26]int{}
		curIndex := left
		for ; curIndex <= right; curIndex++ {
			arr1[s[curIndex]-'a']++
		}
		if arr1 == arr2 {
			res = append(res, left)
		}
		left++
		right++
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
