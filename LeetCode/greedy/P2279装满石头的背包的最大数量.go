package main

import (
	"fmt"
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-bags-with-full-capacity-of-rocks/

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}

	fmt.Println(capacity)

	// 从小到大排序
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] < capacity[j]
	})

	fmt.Println(capacity)

	res := 0
	for res < len(capacity) && additionalRocks > 0 {
		additionalRocks -= capacity[res]
		if additionalRocks >= 0 {
			res++
		}
	}

	return res
}
