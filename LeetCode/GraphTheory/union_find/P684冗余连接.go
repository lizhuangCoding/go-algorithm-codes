package main

// 并查集
// 力扣：https://leetcode.cn/problems/redundant-connection/description/

/*
在一棵树中，边的数量比节点的数量少 1。如果一棵树有 n 个节点，则这棵树有 n−1 条边。这道题中的图在树的基础上多了一条附加的边，因此边的数量也是 n。
树是一个连通且无环的无向图，在树中多了一条附加的边之后就会出现环，因此附加的边即为导致环出现的边。

可以通过并查集寻找附加的边。初始时，每个节点都属于不同的连通分量。遍历每一条边，判断这条边连接的两个顶点是否属于相同的连通分量。
如果两个顶点属于不同的连通分量，则说明在遍历到当前的边之前，这两个顶点之间不连通，因此当前的边不会导致环出现，合并这两个顶点的连通分量。
如果两个顶点属于相同的连通分量，则说明在遍历到当前的边之前，这两个顶点之间已经连通，因此当前的边导致环出现，为附加的边，将当前的边作为答案返回。
*/
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var union func(x, y int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		parent[rootY] = rootX // y 指向 x
	}

	result := make([]int, 2)

	for i := 0; i < len(edges); i++ {
		n1, n2 := edges[i][0], edges[i][1]
		rootN1, rootN2 := find(n1), find(n2)

		// 说明n1和n2的最终父节点是同一个，这里就会形成一个环，说明这两个节点可能是结果值
		if rootN1 == rootN2 {
			result[0] = n1
			result[1] = n2
		}

		union(n1, n2)
	}
	return result
}
