package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_3(t *testing.T) {
	t.Log("3. Longest Substring Without Repeating Characters https://leetcode.com/problems/longest-substring-without-repeating-characters/description/")
	var testcases = []struct {
		text  string
		input string
		want  int
	}{
		{"Example 1:", "abcabcbb", 3},
		{"Example 2:", "bbbbb", 1},
		{"Example 3:", "pwwkew", 3},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}

}
