package main

func mySqrt(x int) int {
	var i uint64
	ux := uint64(x)
	for i = 1; ; i++ {
		if i*i < ux {
			continue
		} else if i*i == ux {
			return int(i)
		} else {
			return int(i - 1)
		}
	}
}
