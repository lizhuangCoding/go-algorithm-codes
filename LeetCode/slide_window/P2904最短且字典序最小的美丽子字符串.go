package main

import (
	"math"
	"strings"
)

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/shortest-and-lexicographically-smallest-beautiful-string/description/

// 两个规则：1.长度；2.字典序

// 时间复杂度：O(n^2)，其中 n 为 s 的长度。
// 空间复杂度：O(n) 或 O(1)。字符串切片需要 O(n) 的空间，Go 除外。
// 注：利用字符串哈希（或者后缀数组等），可以把比较字典序的时间降至 O(nlogn)，这样可以做到 O(nlogn) 的时间复杂度。
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}

	minLen := math.MaxInt64
	result := ""
	demo := ""
	left := 0

	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			k--
		}

		// 例如：100011001，不仅要清除索引为0位置的'1'，还要清除索引为1-3位置的'0'
		for k < 0 || (left < len(s) && s[left] == '0') {
			if s[left] == '1' {
				k++
			}
			left++
		}

		if k == 0 {
			demo = s[left : right+1]
			// 1.长度；2.字典序
			if minLen > right-left+1 || (minLen == right-left+1 && result > demo) {
				minLen = right - left + 1
				result = demo
			}
		}
	}
	return result
}
