package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/** 判断是否是回文链表**/
func isPalindrome(head *ListNode) bool {
	node := []int{}
	for head != nil {
		node = append(node, head.Val)
		head = head.Next
	}
	i, j := 0, len(node) - 1
	for i <= j {
		if node[i] != node[j] {
			return false	
		}
		i++
		j--
	}
	return true
}

// 链表反转
func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}