package main

import "fmt"

func main() {
	var a = []int{1,2,3, 4, 5}
	var b = make([]int, 5, 10)
	fmt.Println(a)
	append(a, 6, 7,8)
	copy(b, a)
	fmt.Println(b)
}
