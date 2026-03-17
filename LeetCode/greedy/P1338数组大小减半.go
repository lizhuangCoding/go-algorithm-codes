package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/reduce-array-size-to-the-half/description/

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	// 存储每个数字出现的次数
	sli := make([]int, 0)
	for _, v := range m {
		sli = append(sli, v)
	}
	// 从大到小排列
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] > sli[j]
	})

	n, demo := len(arr)/2, 0
	res := 0
	for ; res < len(sli) && demo < n; res++ {
		demo += sli[res]
	}
	return res
}
