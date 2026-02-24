package main

// 矩阵
// 力扣：https://leetcode.cn/problems/game-of-life/description/

// 方法一：暴力，使用额外的空间
// 略...

// 方法二：使用额外的状态
func gameOfLife(board [][]int) {
	// 添加新状态：
	// 2:表示这个细胞过去是死的，现在是活的
	// 3:表示这个细胞过去是活的，现在是死的
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			num := gameOfLifeJudge(i, j, board)

			// 活变死
			if board[i][j] == 1 && (num < 2 || num > 3) {
				board[i][j] = 3
			}
			// 死变活
			if board[i][j] == 0 && num == 3 {
				board[i][j] = 2
			}
		}
	}
	// 修改额外的状态
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			}
			if board[i][j] == 3 {
				board[i][j] = 2
			}

		}
	}

}

// 返回这个细胞周围有多少个活细胞
func gameOfLifeJudge(x, y int, board [][]int) int {
	row, col := len(board), len(board[0])
	result := 0

	// %2 是为了兼容 状态2和3
	if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1]%2 == 1 {
		result++
	}
	if x-1 >= 0 && board[x-1][y]%2 == 1 {
		result++
	}
	if x-1 >= 0 && y+1 <= col-1 && board[x-1][y+1]%2 == 1 {
		result++
	}

	if y-1 >= 0 && board[x][y-1]%2 == 1 {
		result++
	}
	if y+1 <= col-1 && board[x][y+1]%2 == 1 {
		result++
	}

	if x+1 <= row-1 && y-1 >= 0 && board[x+1][y-1]%2 == 1 {
		result++
	}
	if x+1 <= row-1 && board[x+1][y]%2 == 1 {
		result++
	}
	if x+1 <= row-1 && y+1 <= col-1 && board[x+1][y+1]%2 == 1 {
		result++
	}
	return result
}
