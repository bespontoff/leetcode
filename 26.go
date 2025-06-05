package main

func removeDuplicates(nums []int) int {
	result := make([]int, 0)
	var current int
	for _, num := range nums {
		if len(result) == 0 {
			current = num
			result = append(result, current)
		} else {
			if num != current {
				result = append(result, num)
				current = num
			} else {
				continue
			}
		}

	}
	return len(result)
}
