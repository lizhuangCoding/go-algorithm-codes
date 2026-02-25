package dfs_bfs

// 图论，bfs
// 力扣：https://leetcode.cn/problems/jump-game-iii/

func canReach(arr []int, start int) bool {
	queue := make([]int, 0) // 存储下标
	m := make(map[int]bool) // 存储已经走过的下标位置

	queue = append(queue, start)

	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]

		if arr[i] == 0 {
			return true
		}

		// 可以向右走
		if i+arr[i] <= len(arr)-1 && !m[i+arr[i]] {
			queue = append(queue, i+arr[i])
			m[i+arr[i]] = true
		}
		// 可以向左走
		if i-arr[i] >= 0 && !m[i-arr[i]] {
			queue = append(queue, i-arr[i])
			m[i-arr[i]] = true
		}

	}

	return false
}
