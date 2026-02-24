package main

// 分组循环，字符串，贪心
// 力扣：https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/description/

func minCost(colors string, neededTime []int) int {
	result := 0

	for i := 0; i < len(colors); i++ {
		start := i

		totalTime := neededTime[start] // 总时间
		maxTime := neededTime[start]   // 一段连续颜色的气球中：剪除某个气球的最长时间

		for ; i < len(colors)-1 && colors[i] == colors[i+1]; i++ {
			totalTime += neededTime[i+1]
			maxTime = minCostMax(maxTime, neededTime[i+1])
		}

		// 如果这段连续颜色的气球的数量 != 1，那么我们就需要把这段连续颜色的气球剪的只剩一个，那么耗费的最小时间 = 总时间 - 剪除一个气球的最大时间
		// 也就是贪心：保留删除成本最高的颜色，并删除其他颜色。这样得到的删除成本一定是最低的。
		if i-start+1 != 1 {
			result += totalTime - maxTime
		}
	}
	return result
}

func minCostMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
