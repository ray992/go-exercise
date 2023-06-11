package main

import "fmt"

func main() {
	//将切片复制到另一个切片
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Println(s2)

	s3 := append([]int{}, s1...)
	fmt.Println(s3)

	s4 := append(s1[:0:0], s1...)
	fmt.Println(s4)

	//删除
	s5 := append(s1[:1], s1[2:]...)
	fmt.Println(s5)
}
