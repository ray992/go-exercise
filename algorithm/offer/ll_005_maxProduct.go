package main

import "fmt"

// 单词长度最大乘积

func maxProduct(words []string) int {
	length := len(words)
	/*var mask []int
	for i := 0; i < length; i++ {
		mask = append(mask, 0)
	}*/
	mask := make([]int, length)
	for i := 0; i < length; i++ {
		curWord := words[i]
		for j := 0; j < len(curWord); j++ {
			mask[i] |= 1 << (curWord[j] - 'a')
		}
	}
	result := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if mask[i]&mask[j] == 0 {
				curMul := len(words[i]) * len(words[j])
				if curMul > result {
					result = curMul
				}
			}
		}
	}
	return result
}

func maxProduct2(words []string) int {
	length := len(words)
	/*var mask []int
	for i := 0; i < length; i++ {
		mask = append(mask, 0)
	}*/
	mask := make([]int, length)
	maskMap := make(map[int]int, length)
	for i := 0; i < length; i++ {
		curWord := words[i]
		for j := 0; j < len(curWord); j++ {
			mask[i] |= 1 << (curWord[j] - 'a')
		}
		v, ok := maskMap[mask[i]]
		if !ok {
			maskMap[mask[i]] = len(curWord)
		} else {
			if len(curWord) > v {
				maskMap[mask[i]] = len(curWord)
			}
		}
	}
	result := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if mask[i]&mask[j] == 0 {
				vi, _ := maskMap[mask[i]]
				vj, _ := maskMap[mask[j]]
				curMul := vi * vj
				if curMul > result {
					result = curMul
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(maxProduct2([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}))
	fmt.Println(maxProduct2([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct2([]string{"a", "aa", "aaa", "aaaa"}))
}
