package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_69(t *testing.T) {
	t.Log("69. Sqrt(x) https://leetcode.com/problems/sqrtx/description/")
	var testcases = []struct {
		text  string
		input int
		want  int
	}{
		{"Example 1:", 4, 2},
		{"Example 2:", 8, 2},
	}

	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {
			result := mySqrt(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}
}
