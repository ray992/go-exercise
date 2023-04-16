package main

import "fmt"

func permute(nums []int) [][]int {

}

func dfs(nums []int, used map[int]bool, res [][]*int, depth int, length int, path []*int) {
	if depth == length {
		res = append(res, path)
		return
	}
	for i := 0; i < length; i++ {
		use, _ := used[i]
		if use {
			continue
		}
		path = append(path, &nums[i])
		used[i] = true
		dfs(nums, used, res, depth+1, length, path)
		used[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2}))
}
