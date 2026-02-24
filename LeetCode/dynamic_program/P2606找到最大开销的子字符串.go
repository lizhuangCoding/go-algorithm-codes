package main

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/find-the-substring-with-maximum-cost/description/

func maximumCostSubstring(s string, chars string, vals []int) int {
	// 每个字母对应的价值
	mapping := [26]int{}
	for i := 0; i < len(mapping); i++ {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}

	res := 0
	f := 0
	for i := 0; i < len(s); i++ {
		f = maxMaximumCostSubstring(f, 0) + mapping[s[i]-'a']
		res = maxMaximumCostSubstring(res, f)
	}
	return res
}

func maxMaximumCostSubstring(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
