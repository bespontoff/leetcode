package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_66(t *testing.T) {
	t.Log("66. Plus One https://leetcode.com/problems/plus-one/description/")
	var testcases = []struct {
		text  string
		input []int
		want  []int
	}{
		{"Example 1:", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Example 2:", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Example 3:", []int{9}, []int{1, 0}},
		{"Example 4:", []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
		{"Example 5:", []int{9, 9}, []int{1, 0, 0}},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := plusOne(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}
