package main

import "fmt"

func main() {
	//引用类型
	var array = []int{2, 3, 4, 5}
	slice := array[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var ss = make([]int, 2, 10)
	ss[0] = 2
	ss[1] = 5
	//ss[3] = 5 错误，数组下标越界
	ss = append(ss, 4, 6, 10, 12, 16, 2, 3, 6, 7, 9, 0, 9, 9, 9) //追加多个元素
	fmt.Println(ss)
	fmt.Println(len(ss))
	fmt.Println(cap(ss))
	fmt.Println(ss == nil)

	var pp []int
	pp = append(pp, 23, 332, 4421) //追加切片
	pp = append(pp, pp...)
	fmt.Println(pp)
	pp = append(pp, ss...)
	fmt.Println(pp)
}
