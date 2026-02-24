package main

import (
	"math"
)

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/triangle/description/

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	// 初始化
	dp[0][0] = triangle[0][0]
	// 初始化第一列
	for i := 1; i < m; i++ {
		dp[i][0] = triangle[i][0] + dp[i-1][0]
		// 初始化每行的最后一个位置
		// end := len(triangle[i])
		// dp[i][end-1] = triangle[i][end-1] + dp[i-1][end-2]

		// 或者
		dp[i][i] = triangle[i][i] + dp[i-1][i-1]
	}

	// for i := 0; i < m-1; i++ {
	// 	for j := 0; j < len(dp[i])-1; j++ {
	// 		dp[i+1][j+1] = minMinimumTotal(dp[i][j], dp[i][j+1]) + triangle[i+1][j+1]
	// 	}
	// }

	// 或者：
	for i := 1; i < m; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = minMinimumTotal(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}

	minRes := math.MaxInt
	for j := 0; j < len(dp[m-1]); j++ {
		minRes = minMinimumTotal(minRes, dp[m-1][j])
	}
	return minRes
}

func minMinimumTotal(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
// 	fmt.Println(minimumTotal([][]int{{-10}})) // -10
// }
