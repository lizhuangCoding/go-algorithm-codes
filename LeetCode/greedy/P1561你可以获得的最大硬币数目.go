package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-number-of-coins-you-can-get/description/

func maxCoins(piles []int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})

	res := 0
	for i := 0; i < len(piles)/3*2; i += 2 {
		res += piles[i+1]
	}

	return res
}
