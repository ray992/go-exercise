package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool  {
	var countMap = make(map[byte]int, 8)
	for i := 0; i < len(magazine); i++ {
		v, ok := countMap[magazine[i]]
		if ok {
			countMap[magazine[i]] = v+1
		}else {
			countMap[magazine[i]] = 1
		}
	}
	for j := 0; j< len(ransomNote); j++ {
		v, ok := countMap[ransomNote[j]]
		if !ok {
			return false
		}
		if v <= 0 {
			return false
		}
		countMap[ransomNote[j]] = v-1
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
