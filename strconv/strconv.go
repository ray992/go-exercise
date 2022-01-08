package main

import (
	"fmt"
	"strconv"
)

func main() {

	//字符串转int类型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("convert fail")
	}else {
		fmt.Println(i1)
	}

	//int转字符串
	a := 123
	convertString := strconv.Itoa(a)
	fmt.Println(convertString)

	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("f"))

	//字符串转成指定类型的值
	fmt.Println(strconv.ParseInt("100", 8, 64))
	fmt.Println(strconv.ParseInt("10000", 2, 64))
	fmt.Println(strconv.ParseInt("10000", 10, 64))

	fmt.Println(strconv.FormatInt(1000, 16))
}

