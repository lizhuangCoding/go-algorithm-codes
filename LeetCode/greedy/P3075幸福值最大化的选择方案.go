package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/maximize-happiness-of-selected-children/description/

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	res := 0
	k2 := 0
	for i := 0; i < len(happiness) && k2+1 <= k; i++ {
		demo := max(happiness[i]-k2, 0)
		res += demo
		k2++
	}

	return int64(res)
}
