package main

import (
	"strconv"
	"strings"
)

var res string

func tree2str(root *TreeNode) string {
	ans := &strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			ans.WriteString(strconv.Itoa(node.Val))
			hasLeft, hasRight := node.Left != nil || node.Right != nil, node.Right != nil
			if hasLeft {
				ans.WriteByte('(')
				dfs(node.Left)
				ans.WriteByte(')')
			}
			if hasRight {
				ans.WriteByte('(')
				dfs(node.Right)
				ans.WriteByte(')')
			}
		}
	}
	dfs(root)
	return ans.String()
}

func main() {

}
