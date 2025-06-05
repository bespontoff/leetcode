package main

import (
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	result := make([]int, 0)
	s := make([]string, 0)

	for _, digit := range digits {
		s = append(s, strconv.Itoa(digit))
	}
	sNum := strings.Join(s, "")
	num, _ := strconv.Atoi(sNum)
	num++
	s = make([]string, 0)
	for _, char := range strings.Split(strconv.Itoa(num), "") {
		d, _ := strconv.Atoi(char)
		result = append(result, d)
	}
	return result
}
