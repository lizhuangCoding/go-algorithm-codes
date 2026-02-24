package main

// 动态规划 一、入门DP 1.2 打家劫舍
// 力扣：https://leetcode.cn/problems/count-number-of-ways-to-place-houses/

func countHousePlacements(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i <= n; i++ {
		// 街道一侧的DP将与斐波那契数列相似（可以列出例子看看）
		dp[i] = (dp[i-1] + dp[i-2]) % 1_000_000_007
	}
	return dp[n] * dp[n] % 1_000_000_007
}
