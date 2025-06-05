package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	if list1 == (&ListNode{}) && list2 == (&ListNode{}) {
		return &ListNode{}
	}
	if list1 == (&ListNode{}) {
		return list2
	}
	if list2 == (&ListNode{}) {
		return list1
	}

	return result
}
