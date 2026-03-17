package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-ice-cream-bars/description/

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	res := 0
	for i := 0; i < len(costs) && coins > 0; i++ {
		coins -= costs[i]
		if coins >= 0 {
			res++
		}
	}

	return res
}
