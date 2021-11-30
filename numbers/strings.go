package main

import (
	"fmt"
)

func main() {
	//只有强制转换，没有隐式转换
	var s1 string = "abc"
	bytess := []byte(s1)
	bytess[0] = 'm'
	fmt.Println(string(bytess))

	var s2  string = "234沙盒小王子"
	var count int8 = 0
	for _, mm := range  s2 {
		fmt.Println(mm)
		if mm == 50 || mm == 51 || mm == 52 {
			continue
		}
		count++
	}
	fmt.Println(count)
}
