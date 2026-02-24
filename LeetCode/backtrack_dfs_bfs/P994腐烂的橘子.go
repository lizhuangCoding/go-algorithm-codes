package main

// 图论，广度遍历
// 力扣：https://leetcode.cn/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
	queue := make([][]int, 0) // 二维存储腐烂橘子的坐标
	time := 0

	// 初始化，存储腐烂的橘子的坐标
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// 广度遍历
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]

			// 向左走
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				queue = append(queue, []int{x, y - 1})
			}
			// 向上走
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				queue = append(queue, []int{x - 1, y})
			}
			// 向右走
			if y+1 < len(grid[0]) && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				queue = append(queue, []int{x, y + 1})
			}
			// 向下走
			if x+1 < len(grid) && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				queue = append(queue, []int{x + 1, y})
			}
		}
		time++
	}

	// 判断是否都被污染了
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	if time == 0 {
		return 0
	}

	// 第一轮不能算进去，要 time-1
	return time - 1
}
