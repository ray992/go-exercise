package main

import "fmt"

func equalPairs(grid [][]int) int {
	res := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		temp := grid[i]
		for k := 0; k < n; k++ {
			col := make([]int, 0)
			for j := 0; j < n; j++ {
				col = append(col, grid[j][k])
			}
			status := true
			for r := 0; r < n; r++ {
				if temp[r] != col[r] {
					status = false
					break
				}
			}
			if status {
				res++
			}
			col = nil
		}
	}
	return res
}
func equalPairs2(grid [][]int) (ans int) {
	cnt := make(map[[200]int]int)
	for _, row := range grid {
		a := [200]int{}
		for i, v := range row {
			a[i] = v
		}
		cnt[a]++
	}
	for j := range grid[0] {
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}
		ans += cnt[a]
	}
	return
}

func main() {
	fmt.Println(equalPairs2([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	fmt.Println(equalPairs2([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}
