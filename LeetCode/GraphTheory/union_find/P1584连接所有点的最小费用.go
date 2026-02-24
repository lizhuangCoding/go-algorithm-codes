package main

import "sort"

// 最小生成树、Kruskal算法、并查集
// 最小生成树（使用 Kruskal 算法来解决最小生成树（MST）问题，具体是计算给定二维平面上的点集连接成一个连通图所需的最小成本）
// 力扣：https://leetcode.cn/problems/min-cost-to-connect-all-points/description/

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	parent := make([]int, n)
	for i := 0; i < len(parent); i++ {
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

	// 合并
	var union func(x, y int) bool
	union = func(x, y int) bool {
		rootX, rootY := find(x), find(y)
		if rootX == rootY {
			return false // 合并失败
		}
		// 把y合并到x上
		parent[rootY] = rootX
		return true
	}

	// 定义一个 edge 结构体，包含边的两个端点 v 和 m 以及边的距离 dist
	type Edge struct{ v, m, dist int }
	edges := make([]Edge, 0)

	// 通过两层循环遍历所有点对，计算每对点之间的曼哈顿距离，并将边信息存储在 edges 切片中。
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			// 用加权边连接每对点，权重是这些点之间的曼哈顿距离
			var dist int

			demo1, demo2 := points[i][0]-points[j][0], points[i][1]-points[j][1]
			if demo1 > 0 {
				dist += demo1
			} else {
				dist += -demo1
			}

			if demo2 > 0 {
				dist += demo2
			} else {
				dist += -demo2
			}
			edges = append(edges, Edge{v: i, m: j, dist: dist})
		}
	}

	// 按照边的距离从小到大进行排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// 判断加入边是否会形成环
	// 合并，如果不能合并说明会形成环，那么就跳过当前
	ans := 0
	// left := n - 1 // left是剪枝操作，没有left也可以得到正确结果

	// Kruskal 算法核心：遍历排序后的边列表，对于每条边，尝试合并边的两个端点所在的集合。
	// 如果合并成功，说明加入这条边不会形成环，将边的距离累加到结果 ans 中，并将剩余需要连接的边数 left 减 1。当 left 为 0 时，说明已经找到了最小生成树，退出循环。
	for i := 0; i < len(edges); i++ {
		if union(edges[i].v, edges[i].m) {
			ans += edges[i].dist
			// left--
		}
		// if left == 0 {
		// 	break
		// }
	}
	return ans
}
