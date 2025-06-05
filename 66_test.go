package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_66(t *testing.T) {
	t.Log("20. Valid Parentheses https://leetcode.com/problems/valid-parentheses/description/")
	var testcases = []struct {
		text  string
		input []int
		want  []int
	}{
		{"Example 1:", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Example 2:", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Example 3:", []int{9}, []int{1, 0}},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := plusOne(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}

}
