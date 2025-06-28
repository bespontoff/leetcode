package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_83(t *testing.T) {
	t.Log("83. Remove Duplicates from Sorted List https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/")
	var testcases = []struct {
		text  string
		input *ListNode
		want  *ListNode
	}{
		{"Example 1:",
			&ListNode{1, &ListNode{1, &ListNode{2, nil}}},
			&ListNode{1, &ListNode{2, nil}}},
		{"Example 2:",
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{"Example 3:",
			&ListNode{1, &ListNode{1, &ListNode{1, nil}}},
			&ListNode{1, nil},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := deleteDuplicates(tc.input)
			assert.EqualExportedValues(t, tc.want, result)
		})
	}
}
