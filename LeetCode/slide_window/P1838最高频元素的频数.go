package main

import (
	"slices"
)

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/frequency-of-the-most-frequent-element/description/

func maxFrequency(nums []int, k int) int {
	slices.Sort(nums)

	maxLen := 1
	left := 0
	demo := 0 // 已经执行+1的次数（不能超过k次）
	for right := 1; right < len(nums); right++ {
		// 把right-2变为right-1，把right-1变为right（right-1变为right的过程中，要计算right-2变为right所需要的值）。
		// （第一次循环：right-2向right-1靠齐，第二次循环：right-1向right靠齐，right-2向right靠齐）
		demo += (nums[right] - nums[right-1]) * (right - left)

		for demo > k {
			demo -= nums[right] - nums[left] // 减去nums[left]增加的值
			left++
		}

		if maxLen < right-left+1 {
			maxLen = right - left + 1
			// fmt.Println(left, right)
		}
	}
	return maxLen
}

// func main() {
// 	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))     // 3
// 	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5)) // 2
// 	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))     // 1
// }
