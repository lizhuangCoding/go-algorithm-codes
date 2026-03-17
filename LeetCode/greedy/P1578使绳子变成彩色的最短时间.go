package main

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/

// 思路：对于每一段连续相同颜色的气球，只需要保留一个（花费最大的），移除其他所有。总花费 = 该段所有气球花费之和 - 该段最大花费
func minCost(colors string, neededTime []int) (ans int) {
	maxT := 0 // 记录当前连续同色段的最大耗时
	for i, t := range neededTime {
		ans += t            // 记录总时间
		maxT = max(maxT, t) // 更新当前段的最大值

		// 判断是否是相同颜色的气球：如果不是，那么就要减去时间
		if i == len(colors)-1 || colors[i] != colors[i+1] {
			// 两种情况：
			// 1. i是最后一个字符 (i == len(colors)-1)
			// 2. 下一个字符颜色不同 (colors[i] != colors[i+1])

			// 减去当前段的最大耗时（保留这个最大的）
			ans -= maxT // 保留耗时最大的气球
			maxT = 0    // 重置，准备计算下一段
		}
	}
	return
}

// func minCost(colors string, neededTime []int) int {
// 	result := 0
//
// 	for i := 0; i < len(colors); i++ {
// 		start := i
//
// 		totalTime := neededTime[start] // 总时间
// 		maxTime := neededTime[start]   // 一段连续颜色的气球中：剪除某个气球的最长时间
//
// 		for ; i < len(colors)-1 && colors[i] == colors[i+1]; i++ {
// 			totalTime += neededTime[i+1]
// 			maxTime = min(maxTime, neededTime[i+1])
// 		}
//
// 		// 如果这段连续颜色的气球的数量 != 1，那么我们就需要把这段连续颜色的气球剪的只剩一个，那么耗费的最小时间 = 总时间 - 剪除一个气球的最大时间
// 		if i-start+1 != 1 {
// 			result += totalTime - maxTime
// 		}
// 	}
// 	return result
// }
