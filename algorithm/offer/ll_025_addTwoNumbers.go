package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/** 链表中的两数相加 **/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var slice1 []int
	var slice2 []int
	for l1 != nil {
		slice1 = append(slice1, l1.Val)
        l1 = l1.Next
	}
	for l2 != nil{
		slice2 = append(slice2, l2.Val)
        l2 = l2.Next
	}
	var slice3 []int
	pre := 0
	p1, p2, cur := 0, 0, 0
	i, j := len(slice1) - 1, len(slice2) - 1;
	for i >=0 || j >= 0{
		if i < 0 {
			p1 = 0
		}else {
			p1 = slice1[i]
		}
		if j < 0 {
			p2 = 0
		}else {
			p2 = slice2[j]
		}
		cur = p1 + p2 + pre
        if pre > 0 {
			pre--
		}
		if cur >= 10 {
			pre++
			cur = cur - 10
		}
		slice3 = append(slice3, cur)
        i--
        j--
	}
    dummy := &ListNode{-1, nil}
	temp := dummy
     if pre > 0 {
		temp.Next = &ListNode{pre, nil}
        temp = temp.Next
	}
	for k := len(slice3) - 1; k >= 0; k-- {
		temp.Next = &ListNode{slice3[k], nil}
		temp = temp.Next
	}
	return dummy.Next
}

func main() {

}