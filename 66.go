package main

func plusOne(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		for i := len(digits) - 1; i > 0; i-- {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				break
			}
			if i == 0 {
				tmp := make([]int, 0)
				tmp = append(tmp, 1)
				digits = append(tmp, digits...)
			}
		}
	} else {
		digits[len(digits)-1]++
	}
	return digits
}
