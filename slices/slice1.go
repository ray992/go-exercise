package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{1,2,3, 4, 5, 6}
	var b = make([]int, 6, 10)
	a = append(a, 6, 7,8)
	fmt.Println(a)
	copy(b, a) //复制
	fmt.Println(b)

	var c = []int{1,2,3, 10, 9,4,5,6}
	c = append(c[:2], c[3:]...) //删除
	fmt.Println(c)
	sort.Ints(c)
	fmt.Println(c)
}
