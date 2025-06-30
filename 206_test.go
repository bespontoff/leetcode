package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_206(t *testing.T) {
	t.Log("206. Reverse Linked List https://leetcode.com/problems/reverse-linked-list/")
	var testcases = []struct {
		text  string
		input *ListNode
		want  *ListNode
	}{
		{"Example 1:",

			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			&ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
		{"Example 2:",
			&ListNode{1, &ListNode{2, nil}},
			&ListNode{2, &ListNode{1, nil}},
		},
		{"Example 3:",
			nil,
			nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := reverseList(tc.input)
			assert.EqualExportedValues(t, tc.want, result)
		})
	}
}
