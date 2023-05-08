package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var Result = &ListNode{}
	current := Result
	first, second := list1, list2

	for first != nil && second != nil {
		if first.Val <= second.Val {
			current.Next = first
			first = first.Next
		} else {
			current.Next = second
			second = second.Next
		}
		current = current.Next

	}
	if first != nil {
		current.Next = first
	}

	if second != nil {
		current.Next = second
	}

	return Result.Next
}
