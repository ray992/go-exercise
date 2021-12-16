package main

import (
	"fmt"
	"unsafe"
)

type myInt int8 //别名
type myArray [3]int8 //别名

type person struct {
	name string
	age int
}

func main() {
	var aa myInt
	fmt.Println(aa)
	var pp myArray
	pp[0] = 1
	fmt.Println(pp)
	var p1 person = person{"ray", 12}
	fmt.Println(p1)

	var p2 person
	fmt.Println(unsafe.Sizeof(p2))

	var ppp struct{}
	fmt.Println(unsafe.Sizeof(ppp))


}
