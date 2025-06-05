package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_21(t *testing.T) {
	t.Log("21. Merge Two Sorted Lists https://leetcode.com/problems/merge-two-sorted-lists/description/")
	var testcases = []struct {
		text  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{"Example 1:", &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{Val: 4}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}},
		{"Example 2:", &ListNode{}, &ListNode{}, &ListNode{}},
		{"Example 3:", &ListNode{}, &ListNode{0, nil}, &ListNode{0, nil}},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			t.Log(tc.list1, tc.list2)
			result := mergeTwoLists(tc.list1, tc.list2)
			assert.EqualExportedValues(t, tc.want, result)
		})
	}
}
