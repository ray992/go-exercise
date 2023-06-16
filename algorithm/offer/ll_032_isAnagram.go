package main

import (
	"fmt"
)

/**
	是否是变位词
**/
func isAnagram(s string, t string) bool {
	n1, n2 := len(s), len(t)
	if n1 != n2  || s == t{
		return false
	}
	countMap := make(map[byte]*[2]int)
	for i := 0; i < n1; i++ {
		b1 := s[i]
		b2 := t[i]
		v, ok := countMap[b1]
		if ok {
			v[0] = v[0] + 1
		}else {
			countMap[b1] = &[2]int{1, 0}
		}

		v1, ok1 := countMap[b2]
		if ok1 {
			v1[1] = v1[1] + 1
		}else {
			countMap[b2] = &[2]int{0, 1}
		}
	}
	for _, v := range countMap {
		if v[0] != v[1] {
			return false	
		}
	}
	return true
}

func main()  {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("a", "a"))
}