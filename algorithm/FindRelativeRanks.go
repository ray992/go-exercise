package main

import (
	"fmt"
	"sort"
	"strconv"
)

//相对名次
func findRelativeRanks(score []int) []string {
	var back = make([]int, len(score), len(score))
	copy(back, score)
	sort.Ints(score)
	var positionMap = make(map[int]int, 10000)
	var position = 1
	for i:= len(score) - 1; i>=0; i-- {
		positionMap[score[i]] = position
		position++
	}
	var res []string
	for _, val := range back{
		order := positionMap[val];
		if order == 1 {
			res = append(res, "Gold Medal")
		}else if order == 2 {
			res = append(res, "Silver Medal")
		}else if order == 3 {
			res = append(res, "Bronze Medal")
		}else {
			res = append(res, strconv.Itoa(order))
		}
	}
	return res
}

//相对名次
func main() {
	var score = []int{10,3,8,9,4}
	res := findRelativeRanks(score)
	fmt.Println(res)
}
