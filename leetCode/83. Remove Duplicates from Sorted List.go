package main

import "fmt"

func main() {
	//lst := &ListNode{
	//	Val:  1,
	//	Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}

	//lst := &ListNode{
	//	Val:  1,
	//	Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	lst := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}}
	//test := &ListNode{Val: 1}
	//fmt.Println(deleteDuplicates(test))
	fmt.Println(deleteDuplicates(lst))
	//fmt.Println(deleteDuplicates(lst))
	//fmt.Println(deleteDuplicates(lst).Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		next := current.Next
		if current.Val == next.Val {
			current.Next = next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
