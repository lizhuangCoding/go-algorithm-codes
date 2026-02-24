package main

// 动态规划 一、入门DP 1.1 爬楼梯
// 力扣：https://leetcode.cn/problems/count-ways-to-build-good-strings/

// 本质是爬楼梯
func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000000 + 7
	res := 0
	dp := make([]int, high+1) // dp[i] 表示构建长度为i的字符串，有dp[i]种
	dp[0] = 1

	for j := 1; j <= high; j++ {
		if zero <= j {
			dp[j] += dp[j-zero]
		}
		if one <= j {
			dp[j] += dp[j-one]
		}
		dp[j] = dp[j] % mod

		if j >= low {
			res += dp[j] % mod
			res %= mod
		}
	}
	return res
}
