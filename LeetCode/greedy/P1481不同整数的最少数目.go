package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/least-number-of-unique-integers-after-k-removals/description/

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	var sli = make([]int, 0)
	for _, v := range m {
		sli = append(sli, v)
	}

	// 对出现次数进行排序，从小到大排
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})

	res := 0
	for i := 0; i < len(sli) && k > 0; i++ {
		k -= sli[i]
		if k >= 0 {
			res++
		}
	}

	return len(sli) - res
}
