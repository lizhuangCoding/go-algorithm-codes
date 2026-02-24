package main

// 动态规划 二、网格图DP 2.2 进阶
// 力扣：https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix/description/

func maxProductPath(grid [][]int) int {
	mod := 1_000_000_007
	m := len(grid)
	n := len(grid[0])
	maxDp := make([][]int, m)
	minDp := make([][]int, m)
	for i := 0; i < m; i++ {
		maxDp[i] = make([]int, n)
		minDp[i] = make([]int, n)
	}

	maxDp[0][0] = grid[0][0]
	minDp[0][0] = grid[0][0]
	// 初始化第一行
	for j := 1; j < n; j++ {
		maxDp[0][j] = maxDp[0][j-1] * grid[0][j]
		minDp[0][j] = maxDp[0][j]
	}
	// 初始化第一列
	for i := 1; i < m; i++ {
		maxDp[i][0] = maxDp[i-1][0] * grid[i][0]
		minDp[i][0] = maxDp[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if grid[i][j] >= 0 {
				maxDp[i][j] = maxMaxProductPath(maxDp[i-1][j], maxDp[i][j-1]) * grid[i][j]
				minDp[i][j] = minMaxProductPath(minDp[i-1][j], minDp[i][j-1]) * grid[i][j] // 负数*正数还是负数
			} else {
				maxDp[i][j] = minMaxProductPath(minDp[i-1][j], minDp[i][j-1]) * grid[i][j]
				minDp[i][j] = maxMaxProductPath(maxDp[i-1][j], maxDp[i][j-1]) * grid[i][j]
			}

		}
	}

	if maxDp[m-1][n-1] < 0 {
		return -1
	}
	return maxDp[m-1][n-1] % mod
}

func maxMaxProductPath(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func minMaxProductPath(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
