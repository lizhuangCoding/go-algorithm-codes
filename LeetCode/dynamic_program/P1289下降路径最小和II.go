package main

import "math"

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/minimum-falling-path-sum-ii/

func minFallingPathSum2(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)

	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 { // 初始化第一行
				dp[i][j] = grid[i][j]
			} else { // 初始化为最大值
				dp[i][j] = math.MaxInt64
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {

			// 找到上一行最小的 dp[][]，并且不能和当前位置为同一列
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				dp[i][j] = minMinFallingPathSum2(dp[i][j], dp[i-1][k]+grid[i][j])
			}

		}
	}

	minRes := math.MaxInt64
	for j := 0; j < n; j++ {
		minRes = minMinFallingPathSum2(minRes, dp[n-1][j])
	}
	return minRes
}

func minMinFallingPathSum2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
