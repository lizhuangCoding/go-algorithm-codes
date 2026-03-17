package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-sum-with-at-most-k-elements/description/

// 每排取最大的 limits[i] 个数，然后再取这些数中最大的 k 个数，求和即为答案。
func maxSum(grid [][]int, limits []int, k int) int64 {
	sli := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		sort.Slice(grid[i], func(x, y int) bool {
			return grid[i][x] > grid[i][y]
		})

		sli = append(sli, grid[i][:limits[i]]...)
	}

	sort.Slice(sli, func(i, j int) bool {
		return sli[i] > sli[j]
	})

	res := 0
	for i := 0; i < k; i++ {
		res += sli[i]
	}
	return int64(res)
}
