package main

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var row []*TreeNode
	row = append(row, root)
	for len(row) > 0 {
		size := len(row)
		result = append(result, row[size-1].Val)
		for i := 0; i < size; i++ {
			curNode := row[i]
			if curNode.Left != nil {
				row = append(row, curNode.Left)
			}
			if curNode.Right != nil {
				row = append(row, curNode.Right)
			}
		}
		row = row[size:]
	}
	return result
}

func main() {

}
