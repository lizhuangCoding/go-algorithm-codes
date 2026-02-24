package main

import "strings"

// 分组循环，字符串
// 力扣：https://leetcode.cn/problems/delete-characters-to-make-fancy-string/

func makeFancyString(s string) string {
	// 因为性能问题，所以使用 strings.Builder，直接使用 string 时间会超限
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		start := i
		for ; i < len(s)-1 && s[i] == s[i+1]; i++ {
		}

		if i-start+1 > 2 {
			// 只添加两个
			result.WriteByte(s[start])
			result.WriteByte(s[start])
		} else {
			// 从 start 添加到 i
			for j := start; j <= i; j++ {
				result.WriteByte(s[j])
			}
		}
	}
	return result.String()
}
