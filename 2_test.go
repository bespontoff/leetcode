package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_2(t *testing.T) {
	t.Log("2. Add Two Numbers https://leetcode.com/problems/add-two-numbers/")
	var testcases = []struct {
		text  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{"Example 1:", &ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
		{"Example 2:", &ListNode{0, nil}, &ListNode{0, nil}, &ListNode{0, nil}},
		{"Example 3:", &ListNode{9, &ListNode{9, nil}}, &ListNode{9, nil}, &ListNode{8, &ListNode{0, &ListNode{1, nil}}}},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			//t.Log(tc.list1, tc.list2)
			result := addTwoNumbers(tc.list1, tc.list2)
			assert.EqualExportedValues(t, tc.want, result)
		})
	}
}
