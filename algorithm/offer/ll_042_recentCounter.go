package main

// 最近请求次数
type RecentCounter struct {
	count []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}


func (this *RecentCounter) Ping(t int) int {
	this.count = append(this.count, t)
	pre := t - 3000
	count := 0
	for i := len(this.count) - 1; i >= 0; i-- {
		if this.count[i] < pre  {
			break
		}
		if this.count[i] >= pre && this.count[i] <= t {
			count ++
		}
	}
	return count
}