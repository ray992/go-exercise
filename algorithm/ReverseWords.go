package main

import (
	"fmt"
	"strings"
)

/**
反转字符串中的单词
*/
func reverseWords(s string) string {
	slice := strings.Split(s, " ")
	var ss = make([]string, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		ss = append(ss, strings.ReplaceAll(slice[i], " ", ""))
	}
	result := ""
	for i := 0; i < len(ss); i++ {
		if ss[i] == "" {
			continue
		}
		result += ss[i]
		if i < len(ss)-1 {
			result += " "
		}
	}
	return strings.Trim(result, " ")
}

func main() {
	fmt.Println("  hello world  ")
	fmt.Println(reverseWords("  hello world  "))
	fields := strings.Fields(" a  b cd e  e")
	fmt.Println(fields)
}
