package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head.Next
	prev := head
	for h.Next != nil {
		if h.Val == prev.Val {
			prev.Next = h.Next
			h = h.Next
			continue
		}
		prev = h
		h = h.Next
	}
	if h.Val == prev.Val {
		prev.Next = nil
	}
	return head
}
