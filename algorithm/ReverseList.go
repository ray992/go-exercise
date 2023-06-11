package main

func reverseList(head *ListNode) *ListNode {
	var num []int
	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}
	root := &ListNode{}
	pre := root
	for i := len(num) - 1; i >= 0; i-- {
		pre.Next = &ListNode{num[i], nil}
		pre = pre.Next
	}
	return root.Next
}

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {

}
