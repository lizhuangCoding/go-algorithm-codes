package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/mice-and-cheese/description/

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	ans := 0
	for i, v := range reward2 {
		ans += v // 先全部给第二只老鼠
		reward1[i] -= v
	}

	sort.Slice(reward1, func(i, j int) bool {
		return reward1[i] > reward1[j]
	})

	for i := 0; i < k; i++ {
		ans += reward1[i]
	}

	return ans
}
