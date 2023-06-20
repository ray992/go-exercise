package main

/** 删除链表的倒数第 n 个结点 **/

/** 计算链表的长度 **/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy *ListNode = &ListNode{0, head}
	ln := getCount(head)
	cur := dummy
	for i := 0; i < ln-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getCount(listNode *ListNode) int {
	n := 0
	for listNode != nil {
		n++
		listNode = listNode.Next
	}
	return n
}

// 双指针法  先走N步
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	firstNode := head
	secondNode := dummy
	for iii := 0; iii < n; iii++ {
		firstNode = firstNode.Next
	}
	for firstNode != nil {
		firstNode = firstNode.Next
		secondNode = secondNode.Next
	}
	secondNode.Next = secondNode.Next.Next
	return dummy.Next
}

func main() {

}
