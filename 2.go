package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	memory := 0
	result := &ListNode{}
	head := result
	for {
		sum := l1.Val + l2.Val + memory
		result.Next = &ListNode{sum % 10, nil}
		result = result.Next
		memory = sum / 10
		if l1.Next == nil || l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if memory == 1 {
		result.Next = &ListNode{memory, nil}
	}
	return head.Next
}
