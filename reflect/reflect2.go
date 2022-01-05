package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}, v1 int64)  {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(v1)
	}
}

func main() {
	var a int64 = 1000
	var b int64 = 9999
	reflectSetValue(&a, b)
	fmt.Println(a)
}
