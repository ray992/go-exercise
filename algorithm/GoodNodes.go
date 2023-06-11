package main

var num = 0

func goodNodes(root *TreeNode) int {
	//good_dfs(root, root.Val)
	return good_dfs2(root, root.Val)
}

func good_dfs2(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	num := 0
	if node.Val >= max {
		num = 1
		max = node.Val
	}
	num += good_dfs2(node.Left, max)
	num += good_dfs2(node.Right, max)
	return num
}

func good_dfs(node *TreeNode, max int) {
	if node == nil {
		return
	}
	if node.Val >= max {
		max = node.Val
		num++
	}
	good_dfs(node.Left, max)
	good_dfs(node.Right, max)
}

func main() {

}
