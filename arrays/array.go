package main

import (
	"fmt"
)

//值类型
func main() {
	var intArray [4]int
	intArray[0] = 1000
	intArray[1] = 999
	fmt.Println(intArray)

	var stringArray = [3]string{"a", "b", "c"}
	fmt.Println(stringArray)

	var string2Array = [3]string{"c", "d", "e"}
	fmt.Println(string2Array)

	var p = [...]int{3, 4, 5}
	fmt.Println(p)
	for i := 0; i < len(p); i++ {
		fmt.Print(p[i], "-")
	}
	fmt.Println()
	for index, val := range string2Array {
		fmt.Print(index, "-", val)
		fmt.Println()
	}
	fmt.Println()
	var intSArray = [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < len(intSArray); i++ {
		for j := 0; j < len(intSArray[i]); j++ {
			fmt.Print(intSArray[i][j])
		}
	}
	fmt.Println()
	for _, val := range intSArray {
		for _, val2 := range val {
			fmt.Print(val2)
		}
	}
}
