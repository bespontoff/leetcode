package main

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	stack := make([]string, 0)
	for _, char := range s {
		switch string(char) {
		case "{", "[", "(":
			stack = append(stack, string(char))
		case "}":
			var value string
			if len(stack) == 0 {
				return false
			}
			stack, value = stack[:len(stack)-1], stack[len(stack)-1]
			if value != "{" {
				return false
			}
		case "]":
			var value string
			if len(stack) == 0 {
				return false
			}
			stack, value = stack[:len(stack)-1], stack[len(stack)-1]
			if value != "[" {
				return false
			}
		case ")":
			var value string
			if len(stack) == 0 {
				return false
			}
			stack, value = stack[:len(stack)-1], stack[len(stack)-1]
			if value != "(" {
				return false
			}

		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
