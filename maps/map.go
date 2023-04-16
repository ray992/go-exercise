package main

import "fmt"

func main() {
	var m = make(map[string]int, 8)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)

	v, ok := m["d"]
	fmt.Println(v, ok)
	v1, ok1 := m["b"]
	fmt.Println(v1, ok1)
	delete(m, "b")
	for key, val := range m {
		fmt.Println(key, val)
	}

	var map_ = make(map[string]interface{}, 8)
	map_["a"] = struct{}{}
	fmt.Println(map_)
}
