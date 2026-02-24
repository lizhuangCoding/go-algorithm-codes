package main

// 签到：模拟

func numRookCaptures(board [][]byte) int {
	I, J := 0, 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				I, J = i, j
				// break
			}
		}
	}

	num := 0
	// 往左延伸
	for j := J; j >= 0; j-- {
		if board[I][j] == 'p' {
			num++
			break
		} else if board[I][j] == 'B' {
			break
		}
	}

	// 往右延伸
	for j := J; j < len(board[0]); j++ {
		if board[I][j] == 'p' {
			num++
			break
		} else if board[I][j] == 'B' {
			break
		}
	}

	// 往上延伸
	for i := I; i >= 0; i-- {
		if board[i][J] == 'p' {
			num++
			break
		} else if board[i][J] == 'B' {
			break
		}
	}

	// 往下延伸
	for i := I; i < len(board); i++ {
		if board[i][J] == 'p' {
			num++
			break
		} else if board[i][J] == 'B' {
			break
		}
	}
	return num
}
