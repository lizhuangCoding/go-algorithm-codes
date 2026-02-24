package main

// 数组，二分查找
// 力扣：https://leetcode.cn/problems/h-index-ii/

// 需要找到最左边的那个边界位置
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		// citations[mid] 是第 mid 篇论文的引用次数
		// n-mid 是从 mid 开始的论文数量
		if citations[mid] >= n-mid { // 满足条件，尝试找更大的 h（即更小的 n-mid）。也就是向左移动，看是否还有更早的位置也满足条件
			right = mid - 1
		} else { // 不满足条件，需要向右移动
			left = mid + 1
		}
	}

	// left 是第一个满足 citations[i] >= n-i 的位置
	// 如果没有满足的，left = n，返回 0
	return n - left
}
