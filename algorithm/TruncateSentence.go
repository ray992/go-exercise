package main

import (
	"fmt"
	"strings"
)
//截断句子
func truncateSentence(s string, k int) string {
	slices := strings.Split(s, " ")
	var res []string
	for i := 0; i <= k-1; i++ {
		res = append(res, slices[i])
	}
	resString := strings.TrimRight(fmt.Sprint(res), "")
	resString = strings.ReplaceAll(resString, "[", "")
	resString = strings.ReplaceAll(resString, "]", "")
	return resString
}

func main() {
	res := truncateSentence("Hello how are you Contestant", 4)
	fmt.Println(res)
	res1 := truncateSentence("What is the solution to this problem", 4)
	fmt.Println(res1)
}
