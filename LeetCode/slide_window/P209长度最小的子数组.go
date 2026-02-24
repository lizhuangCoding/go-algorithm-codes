package main

import "math"

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/minimum-size-subarray-sum/description/

// 方法一：滑动窗口
// 时间复杂度：O(n)。空间复杂度：O(1)
// left 和 right 都最多遍历一次，因此总的遍历次数是 O(n)。（最多遍历次数不超过2*n次）
func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt64
	left, right := 0, 0

	sum := 0
	for ; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			minLen = minMinSubArrayLen(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt64 {
		return 0
	}
	return minLen
}

func minMinSubArrayLen(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 方法二：前缀和+二分法
// 时间复杂度：O(nlogn)。空间复杂度：O(n)
// func minSubArrayLen(s int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
//
// 	ans := math.MaxInt32
// 	// 前缀和
// 	sums := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		sums[i] = sums[i-1] + nums[i-1]
// 	}
//
// 	for i := 1; i <= n; i++ {
// 		target := s + sums[i-1]
// 		// sort.SearchInts 是 Go 标准库提供的二分查找方法，它用于在一个升序数组中找到目标值或者大于目标值的第一个位置。
// 		bound := sort.SearchInts(sums, target)
// 		if bound <= n {
// 			ans = minMinSubArrayLen(ans, bound-(i-1))
// 		}
// 	}
//
// 	if ans == math.MaxInt32 {
// 		return 0
// 	}
// 	return ans
// }
