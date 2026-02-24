package main

import (
	"math"
)

// 字符串
// 力扣：https://leetcode.cn/problems/string-to-integer-atoi/

// 模拟
func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	res := 0

	// 空格
	start := 0
	for ; start < len(s); start++ {
		if s[start] != ' ' {
			break
		}
	}

	// 正负
	isFushu := false
	if start < len(s) && s[start] == '-' {
		isFushu = true
		start++
	} else if start < len(s) && s[start] == '+' {
		start++
	}

	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
		} else {
			break
		}

		if res > math.MaxInt32 || res < math.MinInt32 {
			if isFushu {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

	}

	if isFushu {
		return -1 * res
	}
	return res
}
