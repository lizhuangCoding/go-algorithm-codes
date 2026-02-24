package main

import (
	"slices"
)

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/description/

// 排序+滑动窗口
// 题解：https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/solutions/2345805/pai-xu-shuang-zhi-zhen-by-endlesscheng-hbqx/
// 原问题等价于：排序后，找最长的连续子数组，其最大值减最小值不超过 2k。
func maximumBeauty(nums []int, k int) int {
	maxLen := 0
	left := 0

	// 排序
	slices.Sort(nums)

	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// 方法二：差分数组
// 题解：https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/solutions/2806897/shu-zu-de-zui-da-mei-li-zhi-by-leetcode-9jgs0/
