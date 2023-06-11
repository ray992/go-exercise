package main

import "fmt"

/**
交替字符串
*/
func mergeAlternately(word1 string, word2 string) string {
	var n1 = len(word1)
	var n2 = len(word2)
	ans := make([]byte, 0, n1+n2)
	for i := 0; i < n1 || i < n2; i++ {
		if i < n1 {
			ans = append(ans, word1[i])
		}
		if i < n2 {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(mergeAlternately("abc", "adddd"))
}
