package main

import "math"

// 最大层内元素和
func maxLevelSum(root *TreeNode) int {
	var et []*TreeNode
	et = append(et, root)
	max := -math.MaxInt64
	result := 1
	layer := 1
	for len(et) > 0 {
		size := len(et)
		curVal := 0
		for i := 0; i < size; i++ {
			curVal += et[i].Val
			if et[i].Left != nil {
				et = append(et, et[i].Left)
			}
			if et[i].Right != nil {
				et = append(et, et[i].Right)
			}
		}
		et = et[size:]
		if curVal > max {
			result = layer
			max = curVal
		}
		layer++
	}
	return result
}

func main() {

}
