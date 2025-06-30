package main

func reverseList(head *ListNode) *ListNode {
	prev := head
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
	cur.Next = reverseList(head)
	return cur
}
