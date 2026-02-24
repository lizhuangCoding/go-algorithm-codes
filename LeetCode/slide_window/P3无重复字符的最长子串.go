package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 不要截取字符串，直接用容器，map，set之类的，否则截取字符串很麻烦
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) <= 1 {
// 		return len(s)
// 	}
//
// 	m := make(map[byte]int)
// 	maxLen := 0
// 	left := 0
// 	// 外层是右侧，里层是左侧
// 	for right := 0; right < len(s); right++ {
// 		// 如果有重复的，就先删除，然后添加
// 		for left < right && m[s[right]] != 0 {
// 			delete(m, s[left])
// 			left++
// 		}
// 		m[s[right]]++
//
// 		maxLen = maxLengthOfLongestSubstring(maxLen, len(m))
// 	}
//
// 	return maxLen
// }

func maxLengthOfLongestSubstring(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 更快
func lengthOfLongestSubstring(s string) (ans int) {
	window := [128]bool{} // 也可以用 map，这里为了效率用的数组
	left := 0
	for right, c := range s {
		// 如果窗口内已经包含 c，那么再加入一个 c 会导致窗口内有重复元素
		// 所以要在加入 c 之前，先移出窗口内的 c
		for window[c] { // 窗口内有 c
			window[s[left]] = false
			left++ // 缩小窗口
		}
		window[c] = true                                     // 加入 c
		ans = maxLengthOfLongestSubstring(ans, right-left+1) // 更新窗口长度最大值
	}
	return
}
