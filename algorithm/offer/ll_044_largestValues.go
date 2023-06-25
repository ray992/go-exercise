package main

import "math"

/**
求每层节点最大值
*/

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		layer := len(q)
		max := math.MinInt64
		for i := 0; i < layer; i++ {
			if max < q[i].Val {
				max = q[i].Val
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		result = append(result, max)
		q = q[layer:]
	}
	return result
}

func main() {

}
