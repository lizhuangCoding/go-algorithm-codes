package main

// 模拟
// 力扣：https://leetcode.cn/problems/find-winner-on-a-tic-tac-toe-game/description/

func tictactoe(moves [][]int) string {
	n := len(moves)
	var arr [3][3]string

	for i, v := range moves {
		if i%2 == 0 {
			arr[v[0]][v[1]] = "A"
		} else {
			arr[v[0]][v[1]] = "B"
		}
	}

	// 横
	if arr[0][0] != "" && arr[0][0] == arr[0][1] && arr[0][1] == arr[0][2] {
		return arr[0][0]
	}
	if arr[1][0] != "" && arr[1][0] == arr[1][1] && arr[1][1] == arr[1][2] {
		return arr[1][0]
	}
	if arr[2][0] != "" && arr[2][0] == arr[2][1] && arr[2][1] == arr[2][2] {
		return arr[2][0]
	}

	// 竖
	if arr[0][0] != "" && arr[0][0] == arr[1][0] && arr[1][0] == arr[2][0] {
		return arr[0][0]
	}
	if arr[0][1] != "" && arr[0][1] == arr[1][1] && arr[1][1] == arr[2][1] {
		return arr[0][1]
	}
	if arr[0][2] != "" && arr[0][2] == arr[1][2] && arr[1][2] == arr[2][2] {
		return arr[0][2]
	}

	// 斜
	if arr[0][0] != "" && arr[0][0] == arr[1][1] && arr[1][1] == arr[2][2] {
		return arr[0][0]
	}
	if arr[0][2] != "" && arr[0][2] == arr[1][1] && arr[1][1] == arr[2][0] {
		return arr[0][2]
	}

	if n == 9 {
		return "Draw"
	}
	return "Pending"
}
