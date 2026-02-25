package GraphTheory

// 图论
// 力扣：https://leetcode.cn/problems/find-the-town-judge/description/

// 计算各节点的入度和出度：
// 在法官存在的情况下，法官不相信任何人，每个人（除了法官外）都信任法官，且只有一名法官。因此法官这个节点的入度是 n−1, 出度是 0。
func findJudge(n int, trust [][]int) int {
	// in 被信任，out 信任其他人
	in, out := make([]int, n+1), make([]int, n+1)

	for _, t := range trust {
		in[t[1]]++
		out[t[0]]++
	}

	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
