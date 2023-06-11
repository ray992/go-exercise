package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	m1 := make(map[byte]int, 10000)
	m2 := make(map[byte]int, 10000)
	for i := 0; i < len(word1); i++ {
		_, ok := m1[word1[i]]
		if ok {
			m1[word1[i]] = m1[word1[i]] + 1
		} else {
			m1[word1[i]] = 1
		}
	}
	for i := 0; i < len(word2); i++ {
		_, ok := m2[word2[i]]
		if ok {
			m2[word2[i]] = m2[word2[i]] + 1
		} else {
			m2[word2[i]] = 1
		}
	}
	if len(m1) != len(m2) {
		return false
	}
	s1 := make([]int, len(m1))
	s2 := make([]int, len(m2))
	for key, val := range m1 {
		_, ok := m2[key]
		if !ok {
			return false
		}
		s1 = append(s1, val)
	}
	for key, val := range m2 {
		_, ok := m1[key]
		if !ok {
			return false
		}
		s2 = append(s2, val)
	}
	sort.Ints(s1)
	sort.Ints(s2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("cabbba", "aabbss"))
	fmt.Println(closeStrings("uau", "ssx"))
}
