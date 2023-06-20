package main

type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}

/** 展平多级双向链表**/
func flatten(root *Node) *Node {
	var stack []*Node
	var pre *Node
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur == nil {
			cur = stack[len(stack) - 1]
			stack = stack[0:len(stack) -1]
		}
		if cur.Child != nil {
			if cur.Next != nil {
				stack = append(stack, cur.Next)
			}
			cur.Next = cur.Child
			cur.Child = nil
		}
		cur.Prev = pre
		if pre != nil {
			pre.Next = cur
		}
		pre = cur
		cur =  cur.Next
	}
	return root
}



func main()  {
	
}