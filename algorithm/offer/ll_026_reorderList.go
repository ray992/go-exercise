package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 重排链表 （线性表）
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	i, j := 0, len(nodes) - 1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

func main() {

}
