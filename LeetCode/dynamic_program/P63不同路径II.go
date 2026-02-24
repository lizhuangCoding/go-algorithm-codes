package main

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 枝剪操作：如果在起点或终点出现了障碍，直接返回0
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列
	for i := 0; i < m; i++ {
		// 有障碍物，永远无法到达这里以及后面
		if obstacleGrid[i][0] == 1 {
			break
		} else {
			dp[i][0] = 1
		}
	}
	// 初始化第一行
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果是障碍物，直接跳过这次
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
