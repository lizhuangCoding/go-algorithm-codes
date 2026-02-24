package main

import "sort"

// 二分查找 或 双指针
// 力扣：https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/

// 方法一：暴力
// 省略

// 方法二：二分查找
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
func countPairs(nums []int, target int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		res += sort.SearchInts(nums[0:i], target-nums[i])
	}
	return res
}

// 方法三：双指针
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
// func countPairs(nums []int, target int) int {
// 	sort.Ints(nums)
// 	res := 0
// 	for i, j := 0, len(nums) - 1; i < j; i++ {
// 		for i < j &&  nums[i] + nums[j] >= target {
// 			j--
// 		}
// 		res += j - i
// 	}
// 	return res
// }
