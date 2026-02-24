package main

// 数组
// 力扣：https://leetcode.cn/problems/pascals-triangle/

// 杨辉三角形
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}
