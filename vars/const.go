package main

import "fmt"

const name = "abc"
const age = 27
const (
	d = 12
	b = 10
	c
	ff
	gg
)

func main() {
	age := age + 1
	fmt.Println(name, age)
	fmt.Println(d, b, c, ff, gg)
}
