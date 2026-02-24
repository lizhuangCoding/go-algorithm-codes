package main

// 数组，区间问题
// 力扣：https://leetcode.cn/problems/insert-interval

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	isMerged := false

	for _, interval := range intervals {
		if interval[0] > right { // 在插入区间的右侧且无交集
			if !isMerged {
				ans = append(ans, []int{left, right})
				isMerged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left { // 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else { // 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	if !isMerged {
		ans = append(ans, []int{left, right})
	}

	return
}
