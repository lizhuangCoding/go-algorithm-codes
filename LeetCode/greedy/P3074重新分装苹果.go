package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/apple-redistribution-into-boxes/description/

func minimumBoxes(apple []int, capacity []int) int {
	totalApples := 0
	for i := 0; i < len(apple); i++ {
		totalApples += apple[i]
	}

	result := 0

	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	for i := 0; i < len(capacity) && totalApples > 0; i++ {
		totalApples -= capacity[i]
		result++
	}

	return result
}
