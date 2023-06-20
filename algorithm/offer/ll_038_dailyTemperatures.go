package main

/** 每日温度 **/

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		cur := temperatures[i]
		for len(stack) > 0 && cur > temperatures[stack[len(stack) - 1]] {
			pre := stack[len(stack) - 1]
			stack = stack[0: len(stack) - 1]
			res[pre] = i - pre
		}
		stack = append(stack, i)
	}
	return res

}