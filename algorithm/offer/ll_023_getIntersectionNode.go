package main

/**
	两个链表的第一个重合节点
**/

/*type ListNode struct {
	Val int
	Next *ListNode
}*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	if headA == nil || headB == nil {
		return nil
	}
	for headA != nil || headB != nil {
		if headA != nil {
			_, ok := nodeMap[headA]
			if ok {
				return headA
			}
			nodeMap[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			_, ok2 := nodeMap[headB]
			if ok2 {
				return headB
			}
			nodeMap[headB] = true
			headB = headB.Next
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}
		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeA
}

func main() {

}
