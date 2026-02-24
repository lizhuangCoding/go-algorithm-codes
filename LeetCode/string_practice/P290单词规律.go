package main

import "strings"

// 字符串：哈希表，一一对应
// 力扣：https://leetcode.cn/problems/word-pattern/description/

func wordPattern(pattern string, s string) bool {
	p2s := make(map[string]string)
	s2p := make(map[string]string)
	arr := strings.Split(s, " ")

	if len(pattern) != len(arr) {
		return false
	}

	for i := range pattern {
		x, y := string(pattern[i]), arr[i]

		if (len(p2s[x]) != 0 && p2s[x] != y) || (len(s2p[y]) > 0 && s2p[y] != x) {
			return false
		}
		p2s[x] = y
		s2p[y] = x
	}
	return true
}
