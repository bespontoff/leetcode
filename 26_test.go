package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_26(t *testing.T) {
	t.Log("26. Remove Duplicates from Sorted Array https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/")
	q1 := make([]int, 0)
	q1 = append(q1, 1, 1, 2)

	q2 := make([]int, 0)
	q2 = append(q1, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4)

	var testcases = []struct {
		text  string
		input []int
		want  int
	}{
		{"Example 1:", q1, 2},
		{"Example 2:", q2, 5},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := removeDuplicates(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}

}
