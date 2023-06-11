package main

func deleteMiddle(head *ListNode) *ListNode {
	n := 0
	pre := head
	for pre != nil {
		pre = pre.Next
		n++
	}
	if n == 1 {
		return nil
	}
	mid := n / 2
	i := 0
	pre1 := head
	for i < mid {
		if i+1 == mid {
			if pre1.Next.Next == nil {
				pre1.Next = nil
				break
			}
			pre1.Next = pre1.Next.Next
		} else {
			pre1 = pre1.Next
		}
		i++
	}
	return head
}

func main() {

}
