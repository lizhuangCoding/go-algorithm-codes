package main

// 回溯
// 力扣：https://leetcode.cn/problems/sudoku-solver/description/

func solveSudoku(board [][]byte) {
	solveSudokuBackTrack(board)
}

func solveSudokuBackTrack(board [][]byte) bool {
	for i := 0; i < 9; i++ { // 行
		for j := 0; j < 9; j++ { // 列

			if board[i][j] != '.' {
				continue
			}

			// 填1-9，看一看哪个合适
			for k := '1'; k <= '9'; k++ {
				// 如果合适就添加到数组中，并进行下一个递归
				if solveSudokuIsValid(i, j, byte(k), board) {
					board[i][j] = byte(k)

					// 如果下一个递归返回true，就可以直接退出了
					if solveSudokuBackTrack(board) {
						return true
					}

					// 复原
					board[i][j] = '.'
				}
			}
			// 9个数字都试过了，但是都不符合，就返回false
			return false
		}
	}
	// 遍历结束，说明成功找到了
	return true
}

// 判断放在这个位置上是否合法
// func solveSudokuIsValid(row int, col int, val byte, board [][]byte) bool {
// 	// 判断放在这一行是否重复
// 	for i := 0; i < 9; i++ {
// 		if board[row][i] == val {
// 			return false
// 		}
// 	}
//
// 	// 判断放在这一列是否重复
// 	for i := 0; i < 9; i++ {
// 		if board[i][col] == val {
// 			return false
// 		}
// 	}
//
// 	// 判断这个 3*3 格子内是否重复
// 	// if row == 0 || row == 1 || row == 2 {
// 	// 	row = 0
// 	// } else if row == 3 || row == 4 || row == 5 {
// 	// 	row = 3
// 	// } else if row == 6 || row == 7 || row == 8 {
// 	// 	row = 6
// 	// }
// 	//
// 	// if col == 0 || col == 1 || col == 2 {
// 	// 	col = 0
// 	// } else if col == 3 || col == 4 || col == 5 {
// 	// 	col = 3
// 	// } else if col == 6 || col == 7 || col == 8 {
// 	// 	col = 6
// 	// }
//
// 	// 或者
// 	row = row / 3 * 3
// 	col = col / 3 * 3
//
// 	for i := row; i < row+3; i++ {
// 		for j := 0; j < col+3; j++ {
// 			if board[i][j] == val {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// 使用位运算优化：
// 判断放在这个位置上是否合法
func solveSudokuIsValid(row int, col int, val byte, board [][]byte) bool {
	var rowMask, colMask, boxMask uint16
	for i := 0; i < 9; i++ {
		// 判断放在这一行是否重复
		if board[row][i] != '.' {
			num := board[row][i] - '0'
			rowMask |= 1 << num
		}
		// 判断放在这一列是否重复
		if board[i][col] != '.' {
			num := board[i][col] - '0'
			colMask |= 1 << num
		}

		// 计算当前元素所在的 3x3 子数独区域内的位置。
		boxI := (row/3)*3 + i/3
		boxJ := (col/3)*3 + i%3
		if board[boxI][boxJ] != '.' {
			// 如果该位置有数字，通过 num := board[boxI][boxJ] - '0' 转换为数字，然后 boxMask |= 1 << num 将该数字对应的位置置为 1。
			num := board[boxI][boxJ] - '0'
			boxMask |= 1 << num
		}
	}

	// num := val - '0'：将待填入的数字 val 转换为对应的整数。
	// (rowMask>>num)&1 == 1：将 rowMask 右移 num 位，并检查最低位是否为 1，为 1 表示数字已经在当前行出现，不合法。
	// 同理，(colMask>>num)&1 == 1 检查列，(boxMask>>num)&1 == 1 检查 3x3 子数独区域。
	num := val - '0'
	if ((rowMask>>num)&1 == 1) || ((colMask>>num)&1 == 1) || ((boxMask>>num)&1 == 1) {
		return false
	}

	return true
}
