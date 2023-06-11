package main

import "fmt"

/**
独一无二的出线次数
*/
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, 1000)
	for i := 0; i < len(arr); i++ {
		_, ok := m[arr[i]]
		if ok {
			m[arr[i]] = m[arr[i]] + 1
		} else {
			m[arr[i]] = 1
		}
	}
	m1 := make(map[int]bool, 1000)
	for _, v := range m {
		_, ok := m1[v]
		if ok {
			return false
		}
		m1[v] = true
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
