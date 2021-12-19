package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Animal interface {
	Sayer
	Mover //接口嵌套
}

type dog struct {

}

type cat struct {

}

func (d dog) say()  { //实现Sayer接口
	fmt.Println("汪汪汪")
}

func (d dog) move()  {
	fmt.Println("奔跑")
}

func (c cat) say()  {
	fmt.Println("喵喵喵")
}

func show(a interface{})  {
	fmt.Println(a)
}

func main() {
	var x Sayer
	var y Mover
	a := dog{}
	x = a
	y = a
	x.say()
	y.move()
	b := cat{}
	x = b
	x.say()
	show("abc")

	var yy interface{}
	var s1 string = "ss"
	yy = s1
	v, ok := yy.(string)//类型断言
	if ok {
		fmt.Println(len(v))
	}else {
		fmt.Println("no type")
	}
}
