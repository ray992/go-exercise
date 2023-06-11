package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	var c = make([]string, 0)
	pre := string(chars[0])
	if n > 1 {
		num := 1
		for i := 1; i < n; i++ {
			cur := string(chars[i])
			if cur == pre {
				num++
			} else {
				c = append(c, pre)
				if num > 1 && num < 10 {
					c = append(c, strconv.Itoa(num))
				} else if num >= 10 {
					for num > 0 {
					}

				}
				pre = cur
				num = 1
			}
		}
		c = append(c, pre)
		if num > 1 && num < 10 {
			c = append(c, strconv.Itoa(num))
			return len(c)
		} else if num >= 10 {
			for j := 0; j < len(strconv.Itoa(num)); j++ {
				c = append(c, "1")
			}
		}
		return len(c)
	}
	return 1
}

func main() {
	fmt.Println(compress([]byte{'a'}))
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'd'}))
}
