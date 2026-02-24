package main

import "fmt"

// 图论，bfs
// 力扣：https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/

func nearestExit(maze [][]byte, entrance []int) int {
	queue := make([][]int, 0)
	queue = append(queue, entrance)

	// 封住起始点，避免走回来做无用功，避免可能会一直循环
	maze[entrance[0]][entrance[1]] = '+'

	time := 0
	for len(queue) != 0 {

		l := len(queue)
		for i := 0; i < l; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]

			// 到达了边缘，可以退出了（记得排除第一次走）
			if time != 0 && (x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[0])-1) {
				return time
			}

			// 向左走
			if y-1 >= 0 && maze[x][y-1] == '.' {
				queue = append(queue, []int{x, y - 1})
				maze[x][y-1] = '+'
			}
			// 向上走
			if x-1 >= 0 && maze[x-1][y] == '.' {
				queue = append(queue, []int{x - 1, y})
				maze[x-1][y] = '+'
			}
			// 向右走
			if y+1 < len(maze[0]) && maze[x][y+1] == '.' {
				queue = append(queue, []int{x, y + 1})
				maze[x][y+1] = '+'
			}
			// 向下走
			if x+1 < len(maze) && maze[x+1][y] == '.' {
				queue = append(queue, []int{x + 1, y})
				maze[x+1][y] = '+'
			}
		}
		time++
	}

	return -1
}

func main() {
	// fmt.Println(nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
	fmt.Println(nearestExit([][]byte{{'.', '+'}}, []int{0, 0}))
}
