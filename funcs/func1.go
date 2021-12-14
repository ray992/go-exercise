package main

import "fmt"

type cal func(int, int) int //函数类型

func add(a, b int) int   {
	return  a+b
}

func sub(a, b int) int  {
	return a - b
}

func cal_(a, b int, op func(int , int) int) int  { //函数作为参数
	return op(a, b)
}

func do(op string) (func(int, int) int, error)  { //函数作为返回值
	if op == "+" {
		return add, nil
	}else if op == "-" {
		return sub, nil
	}
	return nil, nil
}

func main() {
	var cal1 cal //函数类型
	cal1 = add
	fmt.Println(cal1(4,1))

	var cal2 cal
	cal2 = sub
	fmt.Println(cal2(10,4))

	res := cal_(100, 20, add)
	fmt.Println(res)

	funRes, _ := do("+")
	fmt.Println(funRes(100, 100))

	mul := func(a, b int) int { //匿名函数
		return a*b
	}
	result := mul(10, 9)
	fmt.Println(result)

	mod := func(x, y int) int{
		return x%y
	}(10, 3) //自执行函数
	fmt.Println(mod)
	
	ff := add_(100, 20) //闭包
	fmt.Println(ff(10))
	fmt.Println(ff(-100))
}

func add_(a, b int) func(int ) int  {
	return func(c int) int {
		return a+b -c
	}
}
