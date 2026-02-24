package main

import "math"

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/minimum-path-cost-in-a-grid/

func minPathCost(grid [][]int, moveCost [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	// 初始化
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[0][j] = grid[0][j] // 初始化第一行
			} else {
				dp[i][j] = math.MaxInt64 // 其他的初始化为最大值
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {

			for k := 0; k < n; k++ {
				dp[i][j] = minMinPathCost(dp[i][j], dp[i-1][k]+grid[i][j]+moveCost[grid[i-1][k]][j])
			}

		}
	}

	minRes := math.MaxInt64
	for j := 0; j < n; j++ {
		minRes = minMinPathCost(minRes, dp[m-1][j])
	}
	return minRes
}

func minMinPathCost(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
