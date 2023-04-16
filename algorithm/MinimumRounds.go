package main

import (
	"fmt"
	"sort"
)

func minimumRounds(tasks []int) int {
	sort.Ints(tasks)
	cnt := 0
	val := tasks[0]
	tempCnt := 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i] == val {
			tempCnt++
		} else {
			if tempCnt == 1 {
				return -1
			} else if tempCnt == 2 {
				cnt++
			} else {
				n := tempCnt % 3
				if n == 0 {
					cnt += tempCnt / 3
				} else {
					cnt += tempCnt / 3
					cnt++
				}
			}
			val = tasks[i]
			tempCnt = 1
		}
	}
	if tempCnt == 1 {
		return -1
	} else if tempCnt == 2 {
		cnt++
	} else {
		n := tempCnt % 3
		if n == 0 {
			cnt += tempCnt / 3
		} else {
			cnt += tempCnt / 3
			cnt++
		}
	}
	return cnt
}

func main() {
	var task = []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	fmt.Println(minimumRounds(task))
	var task1 = []int{2, 3, 3}
	fmt.Println(minimumRounds(task1))
}
