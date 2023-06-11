package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	var res = make([][]int, 2)
	m1 := make(map[int]bool, 1000)
	m2 := make(map[int]bool, 1000)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = true
	}
	for k, _ := range m1 {
		_, ok := m2[k]
		if !ok {
			res[0] = append(res[0], k)
		}
	}
	for k, _ := range m2 {
		_, ok := m1[k]
		if !ok {
			res[1] = append(res[1], k)
		}
	}
	return res
}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}
