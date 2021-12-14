package main

import "fmt"

func main() {
	a := 10
	b := &a
	c :=*b
	fmt.Println(a, b, c)
	d := 20
	modify(d)
	fmt.Println(d)
	modify_(&d)
	fmt.Println(d)

	//引用类型需要分配内存空间，而值类型的声明不需要分配空间

	var p *int
	p = new(int)
	//*p = 100
	fmt.Println(*p) //默认值是类型的零值，返回的是指向类型的指针
}

func modify(x int)  {
	x = 10
}

func modify_(x *int)  {
	*x = 10
}