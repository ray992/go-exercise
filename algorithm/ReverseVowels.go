package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	var b = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[i]
	}
	var m = make(map[byte]bool)
	m['a'] = true
	m['e'] = true
	m['i'] = true
	m['o'] = true
	m['u'] = true
	m['A'] = true
	m['E'] = true
	m['I'] = true
	m['O'] = true
	m['U'] = true

	var index = make([]int, 0, len(b))
	for i := 0; i < len(b); i++ {
		val, ok := m[s[i]]
		if ok && val {
			index = append(index, i)
		}
	}
	head := 0
	tail := len(index) - 1
	for {
		if head >= tail {
			break
		}
		b[index[head]], b[index[tail]] = b[index[tail]], b[index[head]]
		head++
		tail--
	}
	return string(b)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(strings.Contains("aeioe", string([]byte{'a'})))
}
