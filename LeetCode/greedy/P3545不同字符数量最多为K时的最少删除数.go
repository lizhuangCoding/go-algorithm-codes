package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-deletions-for-at-most-k-distinct-characters/

func minDeletion(s string, k int) int {
	var arr [26]int
	for _, v := range s {
		arr[v-'a']++
	}

	// 从小到大排列
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})

	res := 0
	// 只保留最后k个字母：arr[:26-k]
	for i := 0; i < len(arr[:26-k]); i++ {
		res += arr[i]
	}

	return res
}
