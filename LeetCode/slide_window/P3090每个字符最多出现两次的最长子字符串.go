package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/description/

func maximumLengthSubstring(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	demo := 0
	left := 0
	for right := 0; right < len(s); right++ {
		m[s[right]]++

		// 判断每次添加元素之后是否符合题意
		for m[s[right]] > 2 {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}

		demo = right - left + 1
		if maxLen < demo {
			maxLen = demo
		}
	}
	return maxLen
}
