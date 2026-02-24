package main

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化第一列
	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	// 初始化第一行
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			// 从(i-1,j)的位置向下走一格 + (i,j-1)的位置向右走一格
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
