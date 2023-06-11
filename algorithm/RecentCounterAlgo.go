package main

import "fmt"

type RecentCounter struct {
	c []int64
}

func Constructor() RecentCounter {
	return RecentCounter{
		c: make([]int64, 1),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.c = append(this.c, int64(t))
	pre := t - 3000
	cur := t
	n := 0
	for i := len(this.c) - 1; i >= 0; i-- {
		if this.c[i] == 0 {
			break
		}
		if this.c[i] >= int64(pre) && this.c[i] <= int64(cur) {
			n++
		}
	}
	return n
}

func main() {
	cur := Constructor()
	fmt.Println(cur.Ping(1))
	fmt.Println(cur.Ping(100))
	fmt.Println(cur.Ping(3001))
	fmt.Println(cur.Ping(3002))
}
