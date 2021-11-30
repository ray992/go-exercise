package main

import "fmt"

const (
	p1 = iota
	p2
	p3
)

const (
	p4 = 5
	p5 = iota
	p6
	p7
)

const (
	t1 = iota
	t2 = 100
	t3 = iota
	t4
)

func main() {
	fmt.Println(p1, p2, p3)
	fmt.Println(p4, p5, p6, p7)
	fmt.Println(t1, t2, t3, t4)
}
