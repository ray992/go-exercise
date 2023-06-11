package main

import (
	"fmt"
	"unsafe"
)

type myInt int8      //别名
type myArray [3]int8 //别名

type child struct {
	name string
	age  int
}

type person struct {
	name     string
	age      int
	children []child
}

func main() {
	var aa myInt
	fmt.Println(aa)
	var pp myArray
	pp[0] = 1
	fmt.Println(pp)
	var p1 person = person{"ray", 12, []child{{"mike", 12}, {"jenny", 10}}}
	fmt.Println(p1)

	var p2 person
	fmt.Println(unsafe.Sizeof(p2))
	fmt.Println(p2.name == "")

	var ppp struct{}
	fmt.Println(unsafe.Sizeof(ppp))

}
