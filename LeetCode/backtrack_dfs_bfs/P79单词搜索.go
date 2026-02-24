package main

// 回溯
// 力扣：https://leetcode.cn/problems/word-search/description/
// 代码优化：https://leetcode.cn/problems/word-search/solutions/2927294/liang-ge-you-hua-rang-dai-ma-ji-bai-jie-g3mmm/?envType=problem-list-v2&envId=LBZR6IMZ

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	isVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		isVisited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if board[i][j] == word[0] {
				isVisited[i][j] = true
				if existBacktrack(i, j, board, isVisited, 1, word) {
					return true
				}
				isVisited[i][j] = false
			}

		}
	}
	return false
}

// @param        x int
// @param        y int
// @param        board [][]byte
// @param        isVisited [][]bool 该位置是否已经使用过
// @param        l int 应该符合的长度
// @param        word string
// @return       bool
func existBacktrack(x, y int, board [][]byte, isVisited [][]bool, l int, word string) bool {
	if l == len(word) {
		return true
	}

	// 向左
	if x > 0 && board[x-1][y] == word[l] && !isVisited[x-1][y] {
		isVisited[x-1][y] = true
		if existBacktrack(x-1, y, board, isVisited, l+1, word) {
			return true
		}
		isVisited[x-1][y] = false
	}
	// 向上
	if y > 0 && board[x][y-1] == word[l] && !isVisited[x][y-1] {
		isVisited[x][y-1] = true
		if existBacktrack(x, y-1, board, isVisited, l+1, word) {
			return true
		}
		isVisited[x][y-1] = false
	}
	// 向右
	if y < len(isVisited[0])-1 && board[x][y+1] == word[l] && !isVisited[x][y+1] {
		isVisited[x][y+1] = true
		if existBacktrack(x, y+1, board, isVisited, l+1, word) {
			return true
		}
		isVisited[x][y+1] = false
	}

	// 向下
	if x < len(isVisited)-1 && board[x+1][y] == word[l] && !isVisited[x+1][y] {
		isVisited[x+1][y] = true
		if existBacktrack(x+1, y, board, isVisited, l+1, word) {
			return true
		}
		isVisited[x+1][y] = false
	}
	return false
}

// func main() {
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")) // true
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))    // true
// 	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))   // false
// 	fmt.Println(exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))                   // true
// }
