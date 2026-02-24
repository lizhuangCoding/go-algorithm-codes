package main

import (
	"math"
)

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/minimum-falling-path-sum/

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = matrix[0][0]
	// 不能初始化第一列：
	// for i := 1; i < m; i++ {
	// 	dp[i][0] += matrix[i][0] + dp[i-1][0]
	// }
	// 初始化第一行
	for j := 1; j < n; j++ {
		dp[0][j] += matrix[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			// 初始化为正上方的元素
			minDemo := dp[i-1][j]
			if j+1 < n {
				minDemo = minMinFallingPathSum(minDemo, dp[i-1][j+1])
			}
			if j-1 >= 0 {
				minDemo = minMinFallingPathSum(minDemo, dp[i-1][j-1])
			}
			dp[i][j] = minDemo + matrix[i][j]
		}
	}

	// fmt.Println(dp)

	minRes := math.MaxInt64
	for j := 0; j < n; j++ {
		minRes = minMinFallingPathSum(minRes, dp[m-1][j])
	}
	return minRes
}

func minMinFallingPathSum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))                                                // 13                                            //
// 	fmt.Println(minFallingPathSum([][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}})) // -36
// 	fmt.Println(minFallingPathSum([][]int{{82, 81}, {69, 33}}))                                                             // 114
// }
