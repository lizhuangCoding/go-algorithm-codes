package main

import (
	"strings"
)

// 签到：字符串
func toLowerCase(s string) string {
	var str strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str.WriteByte(s[i] + 32)
		} else {
			str.WriteByte(s[i])
		}
	}
	return str.String()
	// 或者：
	// return strings.ToLower(s)
}
