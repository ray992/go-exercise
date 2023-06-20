package main

/*type ListNode struct {
	Val int
	Next *ListNode
}*/

/** 链表中环的入口节点 **/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]interface{})
	for head != nil {
		_, ok := nodeMap[head]
		if !ok {
			nodeMap[head] = true
		} else {
			return head
		}
		head = head.Next
	}
	return nil
}

func main() {

}
