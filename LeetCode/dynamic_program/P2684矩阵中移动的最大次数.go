package main

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/

// 方法一：动态规划

// 一开始写错的原因：应该先遍历列，然后遍历行（如果先遍历行，有些列其实是可以走的，但是没有遍历到）
func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if dp[i][j] == 0 || j+1 == n {
				continue
			}
			// 当前坐标的右上
			if i-1 >= 0 && grid[i-1][j+1] > grid[i][j] {
				dp[i-1][j+1] = dp[i][j] + 1
			}
			// 当前坐标的正右
			if grid[i][j+1] > grid[i][j] {
				dp[i][j+1] = dp[i][j] + 1
			}
			// 右下
			if i+1 < m && grid[i+1][j+1] > grid[i][j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
		}
	}

	maxRes := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// fmt.Print(dp[i][j], " ")
			maxRes = maxMaxMoves(maxRes, dp[i][j])
		}
		// fmt.Println()
	}
	return maxRes - 1
}

func maxMaxMoves(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 或者：
// func maxMoves(grid [][]int) int {
// 	m := len(grid)
// 	n := len(grid[0])
// 	canMove := make([][]bool, m) // canMove[i][j] 表示是否可以移动到这个单元格
// 	for i := 0; i < m; i++ {
// 		canMove[i] = make([]bool, n)
// 		canMove[i][0] = true // 初始化第一列
// 	}
//
// 	maxMove := 0 // 最大移动次数
// 	// 枚举每一列，再枚举每一行
// 	for j := 1; j < n; j++ {
// 		for i := 0; i < m; i++ {
// 			// 枚举三种转移位置
// 			for p := -1; p <= 1; p++ {
// 				newRow := i + p
// 				if newRow < 0 || newRow >= m {
// 					continue // 剔除非法状态
// 				}
// 				// 可移动需满足两个条件
// 				// 1. grid[i][j] > grid[newRow][j-1]
// 				// 2. (j == 1 || canMove[newRow][j-1])
// 				if grid[i][j] > grid[newRow][j-1] && canMove[newRow][j-1] {
// 					canMove[i][j] = true
// 					maxMove = j
// 					break // 一旦满足可转移，即退出判断
// 				}
// 			}
// 		}
// 	}
// 	return maxMove
// }

// 方法二：dfs
// 从第一列的任一单元格 (i,0) 开始递归。枚举往右上/右/右下三个方向走，如果走一步后，没有出界，且格子值大于 grid[i][j]，则可以走，继续递归。
// func maxMoves(grid [][]int) (ans int) {
// 	m, n := len(grid), len(grid[0])
//
// 	var dfs func(int, int)
// 	dfs = func(i, j int) {
// 		ans = max(ans, j)
// 		if ans == n-1 { // ans 已达到最大值
// 			return
// 		}
// 		// 向右上/右/右下走一步
// 		for k := max(i-1, 0); k < min(i+2, m); k++ {
// 			// 如果可以走
// 			if grid[k][j+1] > grid[i][j] {
// 				dfs(k, j+1)
// 			}
// 		}
// 		grid[i][j] = 0// 防止再走一遍
// 	}
//
// 	for i := range grid {
// 		dfs(i, 0) // 从第一列的任一单元格出发
// 	}
// 	return
// }

// 方法三：bfs
// 首先把所有行坐标加入到集合中，作为出发点。然后对其依次遍历，对每一个单元格，找到下一个列的相邻单元格，并判断是否严格大于当前单元格。
// 如果是，说明可以移动到达。把所有可到达的单元格行坐标加到集合中，并用于下一轮的搜索。当到达最后一列或者集合为空，搜索结束，返回矩阵中移动的最大次数。
// func maxMoves(grid [][]int) int {
// 	m, n := len(grid), len(grid[0])
// 	q := make(map[int]bool)
//
// 	// 初始化
// 	for i := 0; i < m; i++ {
// 		q[i] = true
// 	}
//
// 	// 遍历列
// 	for j := 1; j < n; j++ {
// 		demo := make(map[int]bool)
// 		for i := range q { // 遍历已经到达的位置
// 			for i2 := i - 1; i2 <= i+1; i2++ {
// 				if 0 <= i2 && i2 < m && grid[i][j-1] < grid[i2][j] {
// 					demo[i2] = true // 可以到达的位置
// 				}
// 			}
// 		}
// 		q = demo
// 		if len(q) == 0 {
// 			return j - 1
// 		}
// 	}
// 	return n - 1
// }

// func main() {
// 	fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))   // 3
// 	fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))                              // 0
// 	fmt.Println(maxMoves([][]int{{1, 2, 0, 0}, {10, 3, 0, 0}, {10, 10, 4, 0}, {10, 10, 10, 5}})) // 3
// 	fmt.Println(maxMoves([][]int{
// 		{65, 200, 263, 220, 91, 183, 2, 187, 175, 61, 225, 120, 39},
// 		{111, 242, 294, 31, 241, 90, 145, 25, 262, 214, 145, 71, 294},
// 		{152, 25, 240, 69, 279, 238, 222, 9, 137, 277, 8, 143, 143},
// 		{189, 31, 86, 250, 20, 63, 188, 209, 75, 22, 127, 272, 110},
// 		{122, 94, 298, 25, 90, 169, 68, 3, 208, 274, 202, 135, 275},
// 		{205, 20, 171, 90, 70, 272, 280, 138, 142, 151, 80, 122, 130},
// 		{284, 272, 271, 269, 265, 134, 185, 243, 247, 50, 283, 20, 232},
// 		{266, 236, 265, 234, 249, 62, 98, 130, 122, 226, 285, 168, 204},
// 		{231, 24, 256, 101, 142, 28, 268, 82, 111, 63, 115, 13, 144},
// 		{277, 277, 31, 144, 49, 132, 28, 138, 133, 29, 286, 45, 93},
// 		{163, 96, 25, 9, 3, 159, 148, 59, 25, 81, 233, 127, 12},
// 		{127, 38, 31, 209, 300, 256, 15, 43, 74, 64, 73, 141, 200},
// 	})) // 3
// }
