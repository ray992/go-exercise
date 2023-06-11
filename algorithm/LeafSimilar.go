package main

import "fmt"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var save []int
	find(root1, &save) // 引用类型， 传指针过去，才能改变切片的值
	find(root2, &save)
	fmt.Println(save)
	if len(save)&1 != 0 {
		return false
	}
	for i := 0; i < len(save)/2; i++ {
		if save[i] != save[i+len(save)/2] {
			return false
		}
	}
	return true
}

func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	var find2 func(node *TreeNode, save *[]int)
	find2 = func(node *TreeNode, save *[]int) {
		if node != nil && node.Left == nil && node.Right == nil {
			*save = append(*save, node.Val)
			return
		}
		if node != nil && node.Left != nil {
			find2(node.Left, save)
		}
		if node != nil && node.Right != nil {
			find2(node.Right, save)
		}
	}
	var save []int
	find2(root1, &save) // 引用类型， 传指针过去，才能改变切片的值
	find2(root2, &save)
	fmt.Println(save)
	if len(save)&1 != 0 {
		return false
	}
	for i := 0; i < len(save)/2; i++ {
		if save[i] != save[i+len(save)/2] {
			return false
		}
	}
	return true
}

func find(node *TreeNode, save *[]int) {
	if node != nil && node.Left == nil && node.Right == nil {
		*save = append(*save, node.Val)
		return
	}
	if node != nil && node.Left != nil {
		find(node.Left, save)
	}
	if node != nil && node.Right != nil {
		find(node.Right, save)
	}
}

func main() {

}
