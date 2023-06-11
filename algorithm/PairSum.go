package main

func pairSum(head *ListNode) int {
	n := 1
	var num []int
	var temp = &ListNode{}
	for head != nil {
		num = append(num, head.Val)
		head = head.Next
		temp.Val = head.Val
		temp.Next = &ListNode{}
	}
	root := &ListNode{}
	pre := root
	for i := len(num) - 1; i >= 0; i-- {
		pre.Next = &ListNode{num[i], nil}
		pre = pre.Next
	}
	reverseHead := root.Next
	p := n>>1 - 1
	max := 0
	for i := 0; i <= p; i++ {
		l1 := temp.Val
		l2 := reverseHead.Val
		sum := l1 + l2
		if sum > max {
			max = sum
		}
		head = temp.Next
		reverseHead = reverseHead.Next
	}
	return max
}

func main() {

}
