package main

// 矩阵
// 力扣：https://leetcode.cn/problems/set-matrix-zeroes/description/

// 方法一：用 O(m+n)额外空间
// func setZeroes(matrix [][]int) {
// 	m := len(matrix)
// 	n := len(matrix[0])
//
// 	// 使用标记数组
// 	rows := make([]bool, m)
// 	cols := make([]bool, n)
//
// 	// 记录标记数组
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == 0 {
// 				rows[i] = true
// 				cols[j] = true
// 			}
// 		}
// 	}
//
// 	// 将标记的行置0
// 	for i := 0; i < m; i++ {
// 		if rows[i] == true {
// 			for j := 0; j < n; j++ {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
//
// 	// 将标记的列置0
// 	for j := 0; j < n; j++ {
// 		if cols[j] == true {
// 			for i := 0; i < m; i++ {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// }

// 方法二：用O(1)空间
// 关键思想: 用matrix第一行和第一列记录该行该列是否有0,作为标志位
// 但是对于第一行,和第一列要设置一个标志位,为了防止自己这一行(一列)也有0的情况.
func setZeroes(matrix [][]int) {
	isRow, isCol := false, false
	row, col := len(matrix), len(matrix[0])

	// 判断第一行是否有0，如果有0，需要把该行都置为0
	for j := 0; j < col; j++ {
		if matrix[0][j] == 0 {
			isRow = true
		}
	}

	// 判断第一列是否有0，如果有0，需要把该列都置为0
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
	}

	// 用第一行和第一列记录该行该列是否有0,作为标志位
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 置0（该操作要在下面的两个操作之前完成）
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 是否需要把第一行变为0
	if isRow {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}

	// 是否需要把第一列变为0
	if isCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}

}
