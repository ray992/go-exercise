package main

import "fmt"

var global int = 100

func sum(a int, b ...int) int {
	fmt.Println(a, b)
	sum := a
	for i := 0; i < len(b); i++ {
		sum+=b[i]
	}
	return sum
}

func printVal()  {
	var global  = 20 //优先访问局部变量
	fmt.Println(global)
}

func main() {
	res := sum(10, 1,2,3)
	fmt.Println(res)
	printVal()
}
