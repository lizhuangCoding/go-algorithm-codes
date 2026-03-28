package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximize-points-after-choosing-k-tasks/description/

func maxPoints(technique1 []int, technique2 []int, k int) int64 {
	res := 0
	for i, v := range technique2 {
		res += v
		technique1[i] -= v
	}

	sort.Slice(technique1, func(i, j int) bool {
		return technique1[i] > technique1[j]
	})

	for i := 0; i < len(technique1); i++ {
		if i < k || technique1[i] > 0 {
			res += technique1[i]
		}
	}

	return int64(res)
}
