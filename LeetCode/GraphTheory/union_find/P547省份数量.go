package main

// 并查集，dfs，bfs
// 力扣：https://leetcode.cn/problems/number-of-provinces/description/

// 方法一：并查集
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// 查找
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// 合并（把y合并到x上）
	var union func(x, y int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)

		parent[rootY] = parent[rootX]
	}

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	ans := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == i {
			ans++
		}
	}
	return ans
}

// 方法二：dfs
// func findCircleNum(isConnected [][]int) (ans int) {
// 	vis := make([]bool, len(isConnected))
//
// 	var dfs func(int)
// 	dfs = func(from int) {
// 		vis[from] = true
// 		// 遍历 form 的一行数据（也就是直接和form相连的元素），然后在对这些直接相连的元素进行递归
// 		for to, conn := range isConnected[from] {
// 			if conn == 1 && !vis[to] {
// 				dfs(to)
// 			}
// 		}
// 	}
//
// 	for i, v := range vis {
// 		if !v {
// 			ans++
// 			dfs(i)
// 		}
// 	}
// 	return
// }

// 方法三：dfs
// func findCircleNum(isConnected [][]int) (ans int) {
// 	vis := make([]bool, len(isConnected))
//
// 	for i, v := range vis {
// 		if !v {
// 			ans++
// 			queue := []int{i}
// 			for len(queue) > 0 {
// 				from := queue[0]
// 				queue = queue[1:]
// 				vis[from] = true
// 				for to, conn := range isConnected[from] {
// 					if conn == 1 && !vis[to] {
// 						queue = append(queue, to)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return
// }
