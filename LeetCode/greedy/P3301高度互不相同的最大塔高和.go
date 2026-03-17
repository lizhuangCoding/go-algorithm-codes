package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximize-the-total-height-of-unique-towers/description/

func maximumTotalSum(maximumHeight []int) int64 {
	sort.Slice(maximumHeight, func(i, j int) bool {
		return maximumHeight[i] > maximumHeight[j]
	})

	res := maximumHeight[0]

	for i := 1; i < len(maximumHeight); i++ {
		maximumHeight[i] = min(maximumHeight[i], maximumHeight[i-1]-1) // 取最小值（maximumHeight[i-1]-1，保证唯一）
		if maximumHeight[i] <= 0 {
			return -1
		}
		res += maximumHeight[i]
	}

	return int64(res)
}
