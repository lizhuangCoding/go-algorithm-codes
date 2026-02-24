package GraphTheory

// dfs。有向边，不是无向边。
// 力扣：https://leetcode.cn/problems/detonate-the-maximum-bombs/
// 题解：https://leetcode.cn/problems/detonate-the-maximum-bombs/solutions/1152450/jian-tu-bao-li-mei-ju-suo-you-qi-dian-by-h4mj/
// 问：为什么不能用并查集？
// 答：注意本题是有向图。例如炸弹 0 可以引爆炸弹 2，炸弹 1 可以引爆炸弹 2，对应有向边 0→2,1→2，那么正确答案是 2。
// 如果用并查集做的话，会把 0,1,2 三个点合并起来，计算出错误的答案 3。

// dfs
// 时间复杂度：O(n^3)，其中 n 是 bombs 的长度。注意图中至多有 O(n^2) 条边，每次 DFS 的时间复杂度为 O(n^2)。
// 空间复杂度：O(n^2)
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	// 找到可以引爆的点
	g := make([][]int, n)
	for i, p := range bombs {
		x, y, r := p[0], p[1], p[2]
		for j, q := range bombs {
			dx := x - q[0]
			dy := y - q[1]
			if j != i && dx*dx+dy*dy <= r*r {
				g[i] = append(g[i], j) // i 可以引爆 j
			}
		}
	}

	isVisited := make([]bool, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		isVisited[i] = true
		cnt := 1
		for j := 0; j < len(g[i]); j++ {
			if !isVisited[g[i][j]] {
				cnt += dfs(g[i][j])
			}
		}
		return cnt
	}

	ans := 0
	for i := 0; i < len(g); i++ {
		// 清空数组
		isVisited = make([]bool, len(bombs))

		demo := dfs(i)

		if ans < demo {
			ans = demo
		}

		// 剪枝操作
		if ans == n {
			break
		}
	}
	return ans
}
