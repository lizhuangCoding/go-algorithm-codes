package main

// 前缀和
// 力扣：https://leetcode.cn/problems/find-the-highest-altitude/

func largestAltitude(gain []int) int {
	prefix := 0
	result := 0

	for i := 0; i < len(gain); i++ {
		prefix += gain[i]
		result = max(result, prefix)
	}

	return result
}
