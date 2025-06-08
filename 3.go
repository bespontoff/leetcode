package main

import "sync"

func lengthOfLongestSubstring(s string) int {
	longest := 0
	wg := sync.WaitGroup{}
	for subLength := 1; subLength <= len(s); subLength++ {
		for begin := 0; begin+subLength <= len(s); begin++ {
			end := begin + subLength
			wg.Add(1)
			go func(sub string) {
				m := make(map[int32]struct{})
				uniq := true
				for _, char := range sub {
					if _, ok := m[char]; !ok {
						m[char] = struct{}{}
					} else {
						uniq = false
						break
					}
				}
				if uniq && len(sub) > longest {
					longest = len(sub)
				}
				wg.Done()
			}(s[begin:end])
		}
	}
	wg.Wait()
	return longest
}
