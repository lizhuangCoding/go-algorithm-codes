package main

// 分组循环，字符串
// 力扣：https://leetcode.cn/problems/consecutive-characters/description/

func maxPower(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		start := i
		for ; i < len(s)-1 && s[i] == s[i+1]; i++ {
		}
		result = maxPowerMax(result, i-start+1)
	}
	return result
}

func maxPowerMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
