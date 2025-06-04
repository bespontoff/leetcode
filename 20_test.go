package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_20(t *testing.T) {
	t.Log("20. Valid Parentheses https://leetcode.com/problems/valid-parentheses/description/")
	var testcases = []struct {
		text  string
		input string
		want  bool
	}{
		{"Example 1:", "()", true},
		{"Example 2:", "()[]{}", true},
		{"Example 3:", "(]", false},
		{"Example 4:", "([])", true},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := isValid(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}

}
