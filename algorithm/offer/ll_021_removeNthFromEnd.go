package main

type ListNode struct {
	Val int
	Next *ListNode
}

/** 删除链表的倒数第 n 个结点 **/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy *ListNode  = &ListNode{0, head}
	ln := getCount(head)
	cur := dummy
	for i:=0; i < ln - n; i ++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getCount(listNode *ListNode) int{
	n := 0
	for listNode != nil {
		n++
		listNode = listNode.Next
	}
	return n
}

func main() {

}
