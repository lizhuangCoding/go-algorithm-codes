package main

// 矩阵，遍历
// 力扣：https://leetcode.cn/problems/valid-sudoku/description/

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' && !isValidSudokuJudge(i, j, len(board), len(board[i]), board) {
				return false
			}
		}
	}
	return true
}

func isValidSudokuJudge(x, y, row, col int, board [][]byte) bool {
	// 判断该行是否有重复元素
	for j := 0; j < col; j++ {
		if j != y && board[x][j] == board[x][y] { // 有重复
			return false
		}
	}

	// 判断该列是否有重复元素
	for i := 0; i < row; i++ {
		if i != x && board[i][y] == board[x][y] { // 有重复
			return false
		}
	}

	// 判断3x3小格子
	// 找出每一个3x3小格子开头的位置
	nX := x / 3 * 3
	nY := y / 3 * 3

	for i := nX; i < nX+3; i++ {
		for j := nY; j < nY+3; j++ {
			if (x != i && y != j) && board[i][j] == board[x][y] {
				return false
			}
		}
	}
	return true
}
