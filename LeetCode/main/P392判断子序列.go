package main

// 字符串，动态规划
// 力扣：https://leetcode.cn/problems/is-subsequence/description/

// 方法一：双指针
func isSubsequence(s string, t string) bool {
	slowIndex, fastIndex := 0, 0
	for ; fastIndex < len(t); fastIndex++ {
		if slowIndex == len(s) {
			break
		}

		if s[slowIndex] == t[fastIndex] {
			slowIndex++
		}
	}

	return slowIndex == len(s)
}

// 方法二：动态规划
// ...
