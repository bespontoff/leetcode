package main

func lengthOfLongestSubstring(s string) int {
	subs := make(map[string]struct{})
	for subLength := 1; subLength <= len(s); subLength++ {
		for begin := 0; begin+subLength <= len(s); begin++ {
			end := begin + subLength
			if _, ok := subs[s[begin:end]]; !ok {
				subs[s[begin:end]] = struct{}{}
			}
		}
	}
	longest := 0
	for sub := range subs {
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
	}
	return longest
}
