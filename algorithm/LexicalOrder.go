package main

import (
	"fmt"
	"sort"
	"strconv"
)

func lexicalOrder(n int) []int {
	var valueString []string
	for i := 0; i < n; i++ {
		ss := strconv.Itoa(i + 1)
		valueString = append(valueString, ss)
	}
	slice := valueString[:]
	sort.Strings(slice)
	var res []int
	for i := 0; i < len(slice); i++ {
		val, _ := strconv.Atoi(slice[i])
		res = append(res, val)
	}
	return res
}

func main() {
	fmt.Println(lexicalOrder(192))
}
