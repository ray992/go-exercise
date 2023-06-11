package main

import (
	"fmt"
	"strconv"
)

// 二进制加法

func addBinary2(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	var result []byte
	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			val, _ := strconv.Atoi(string(a[i]))
			sum += val
			i--
		}
		if j >= 0 {
			val, _ := strconv.Atoi(string(b[j]))
			sum += val
			j--
		}
		result = append(result, strconv.Itoa(sum % 2)[0])
		carry = sum / 2
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func addBinary(a string, b string) string {
	var result []byte
	pre := 0
	i := len(a) - 1
	j := len(b) - 1
	for {
		c1 := a[i]
		c2 := b[j]
		if c1 == '0' {
			if c2 == '0' {
				if pre == 0 {
					result = append(result, '0')
				} else if pre == 1 {
					result = append(result, '1')
					pre = 0
				}
			} else if c2 == '1' {
				if pre == 0 {
					result = append(result, '1')
				} else if pre == 1 {
					result = append(result, '0')
				}
			}
		} else {
			if c2 == '0' {
				if pre == 0 {
					result = append(result, '1')
				} else if pre == 1 {
					result = append(result, '0')
					pre = 1
				}
			} else if c2 == '1' {
				if pre == 0 {
					result = append(result, '0')
					pre = 1
				} else if pre == 1 {
					result = append(result, '1')
					pre = 1
				}
			}
		}
		i--
		j--
		if i < 0 || j < 0 {
			break
		}
	}
	if i >= 0 {
		for i >= 0 {
			if a[i] == '0' {
				if pre == 0 {
					result = append(result, '0')
				} else {
					result = append(result, '1')
					pre = 0
				}
			} else if a[i] == '1' {
				if pre == 0 {
					result = append(result, '1')
					pre = 0
				} else {
					result = append(result, '0')
				}
			}
			i--
		}
	}
	if j >= 0 {
		for j >= 0 {
			if b[j] == '0' {
				if pre == 0 {
					result = append(result, '0')
				} else {
					result = append(result, '1')
					pre = 0
				}
			} else if b[j] == '1' {
				if pre == 0 {
					result = append(result, '1')
					pre = 0
				} else {
					result = append(result, '0')
				}
			}
			j--
		}
	}
	if pre == 1 {
		result = append(result, '1')
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func main() {
	//fmt.Println(addBinary("11", "10"))
	//fmt.Println(addBinary("1", "111"))
	//fmt.Println(addBinary("1", "111"))
	fmt.Println(addBinary2("110010", "10111"))
	// 110010
	//  10111
	//  1001001
}
