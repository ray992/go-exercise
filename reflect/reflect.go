package main

import (
	"fmt"
	"reflect"
)

//反射，获取类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("%v\n", v)
}

//反射，获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Float32:
	case reflect.Float64:
		fmt.Println(v.Float())
	case reflect.Bool:
		fmt.Println(v.Bool())
	case reflect.Complex128:
		fmt.Println(v.Complex())

	}
}

func main() {
	var a float64 = 0.54
	reflectType(a)
	var v float64 = -0.92
	reflectValue(v)
	var v1 string = "abc"
	reflectValue(v1)
}
