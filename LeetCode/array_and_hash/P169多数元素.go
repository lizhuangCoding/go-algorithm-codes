package main

// 数组
// 力扣：https://leetcode.cn/problems/majority-element/description/

// 方法一：暴力
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	for _, v := range nums {
// 		m[v]++
// 	}
// 	n := len(nums) / 2
// 	for k, v := range m {
// 		if v > n {
// 			return k
// 		}
// 	}
// 	return 0
// }

// 方法二：排序（排序之后直接选择数组中间位置的元素，代码略）
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)

// 方法三：
// 。。。
