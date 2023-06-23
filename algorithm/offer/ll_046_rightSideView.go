package main

/**
二叉树的右视图
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height == len(res) {
			res = append(res, node.Val)
		}
		height++
		dfs(node.Right, height)
		dfs(node.Left, height)
	}
	dfs(root, 0)
	return res
}

func main() {

}
