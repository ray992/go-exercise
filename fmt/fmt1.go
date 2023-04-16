package main

import "fmt"

func main() {
	s1 := fmt.Sprint("abc")
	s2 := fmt.Sprintf("my name is %s and age is %d", "mike", 18)
	fmt.Println(s1, s2)

	error := fmt.Errorf("this is a error")
	fmt.Println(error)

	fmt.Printf("%b", 123) //二进制
	fmt.Println()
	fmt.Printf("%c", 123) //unicode
	fmt.Println()
	fmt.Printf("%o", 123) //八进制

	fmt.Println()

	var (
		name string
		age  int
	)
	fmt.Scan(&name, &age)
	fmt.Println(name, age)
}
