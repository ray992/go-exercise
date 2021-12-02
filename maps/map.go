package main

import "fmt"

func main() {
	var m = make(map[string]int , 8)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)

	v, ok := m["c"]
	fmt.Println(v, ok)
	v1, ok1 := m["b"]
	fmt.Println(v1, ok1)
	delete(m, "b")
	for key, val := range m{
		fmt.Println(key, val)
	}
}
