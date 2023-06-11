package main

import "fmt"

func maxVowels(s string, k int) int {
	var mapV = make(map[byte]bool)
	mapV['a'] = true
	mapV['e'] = true
	mapV['i'] = true
	mapV['o'] = true
	mapV['u'] = true
	n := len(s)
	i, j := 0, 0
	max := 0
	count := 0
	_, ok := mapV[s[i]]
	if ok {
		count++
	}
	for {
		if j-i+1 < k {
			j++
			if j == n {
				break
			}
			_, ok := mapV[s[j]]
			if ok {
				count++
			}
		} else {
			if max < count {
				max = count
			}
			_, ok := mapV[s[i]]
			if ok {
				count--
			}
			i++
			if i == n {
				break
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
	fmt.Println(maxVowels("rhythms", 4))
	fmt.Println(maxVowels("tryhard", 4))
}
