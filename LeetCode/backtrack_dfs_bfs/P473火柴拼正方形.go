package main

import (
	"sort"
)

// 回溯
// 力扣：https://leetcode.cn/problems/matchsticks-to-square/description/

// 时间超限（代码写的不太好）
// func makesquare(matchsticks []int) bool {
// 	total := 0
// 	for _, v := range matchsticks {
// 		total += v
// 	}
//
// 	d := make([]int, 4)
//
// 	if total%4 != 0 {
// 		return false
// 	}
//
// 	width := total / 4
// 	for _, v := range matchsticks {
// 		if v > width {
// 			return false
// 		}
// 	}
//
// 	return makesquareBacktrack(matchsticks, 0, d, width)
// }
//
// func makesquareBacktrack(matchsticks []int, index int, d []int, width int) bool {
// 	if index == len(matchsticks) {
// 		for i := 0; i < 4; i++ {
// 			if d[i] != width {
// 				return false
// 			}
// 		}
// 		return true
// 	}
//
// 	for i := index; i < len(matchsticks); i++ {
// 		for j := 0; j < len(d); j++ {
// 			if d[j]+matchsticks[i] <= width {
// 				d[j] += matchsticks[i]
// 				if makesquareBacktrack(matchsticks, i+1, d, width) {
// 					return true
// 				}
// 				d[j] -= matchsticks[i]
// 			}
// 		}
// 	}
// 	return false
// }

// 优化：
// 题解：https://leetcode.cn/problems/matchsticks-to-square/solutions/1528435/huo-chai-pin-zheng-fang-xing-by-leetcode-szdp/
func makesquare(matchsticks []int) bool {
	total := 0
	for _, l := range matchsticks {
		total += l
	}
	if total%4 != 0 {
		return false
	}

	// 降序排列，先分配长度长的元素，最后分配长度短的元素，因为长度短的元素更容易分配到对象
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	edges := [4]int{}

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}

		for i := 0; i < len(edges); i++ {
			// 剪枝操作（避免重复性的操作）
			if i > 0 && edges[i] == edges[i-1] {
				continue
			}

			edges[i] += matchsticks[idx]
			if edges[i] <= total/4 && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}

	return dfs(0)
}
