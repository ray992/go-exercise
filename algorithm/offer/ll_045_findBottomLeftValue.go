package main

/**
二叉树最底层最左边的值
*/
// BFS
func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	res := 0
	for len(q) > 0 {
		layer := len(q)
		res = q[0].Val
		for i := 0; i < layer; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[layer:]
	}
	return res
}

// DFS
func findBottomLeftValue2(root *TreeNode) int {
	curHeight := 0
	val := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			val = node.Val
		}
	}
	dfs(root, 0)
	return val
}
