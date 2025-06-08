package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for list1 != nil && list2 != nil {
		if list1 == nil {
			result.insert(list2)
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			result.insert(list1)
			list1 = list1.Next
			continue
		}
		if list1.Val == list2.Val {
			result.insert(list1)
			list1 = list1.Next
			result.insert(list2)
			list2 = list2.Next
		} else if list1.Val < list2.Val {
			result.insert(list1)
			list1 = list1.Next
		} else {
			result.insert(list2)
			list2 = list2.Next
		}

	}
	return result
}

func (ln *ListNode) insert(n *ListNode) {
	current := ln
	for current != nil {
		current = current.Next
	}
	current = n
}
