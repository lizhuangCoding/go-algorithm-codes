package main

// 矩阵
// 力扣：https://leetcode.cn/problems/matrix-diagonal-sum/description/

func diagonalSum(mat [][]int) int {
	n := len(mat)
	res := 0

	// 主对角线
	for i := 0; i < n; i++ {
		res += mat[i][i]
	}
	// 副对角线
	i, j := 0, n-1
	for i != n {
		res += mat[i][j]
		i++
		j--
	}

	// n如果为奇数会有一个数重叠
	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}

// 或者：
// func diagonalSum(mat [][]int) int {
// 	n := len(mat)
// 	sum := 0
// 	mid := n / 2
// 	for i := 0; i < n; i += 1 {
// 		sum += mat[i][i] + mat[i][n - 1 - i]
// 	}
// 	if n & 1 == 1 {
// 		sum -= mat[mid][mid]
// 	}
// 	return sum
// }
