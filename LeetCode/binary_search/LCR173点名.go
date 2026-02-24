package main

// 二分查找
// 力扣：https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/description/

// 方法一：位运算
// func takeAttendance(records []int) int {
// 	res := 0
// 	for i := 0; i <= len(records); i++ {
// 		res ^= i
// 	}
// 	for i := 0; i < len(records); i++ {
// 		res ^= records[i]
// 	}
// 	return res
// }

// 方法二：求和，等差数列
// func takeAttendance(records []int) int {
// 	n := len(records)
// 	sum := n * (n + 1) / 2
// 	for i := 0; i < n; i++ {
// 		sum -= records[i]
// 	}
// 	return sum
// }

// 方法三：暴力
// func takeAttendance(records []int) int {
// 	n := len(records)
// 	for i := 0; i < n; i++ {
// 		if records[i] != i {
// 			return i
// 		}
// 	}
// 	return -1
// }

// 方法四：二分查找
func takeAttendance(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == mid { // 如果 nums[mid] 等于 mid，说明这个区间内没有缺失的数
			left = mid + 1
		} else { // 如果 nums[mid] 不等于 mid，缺失的数就在 [start, mid] 区间
			right = mid
		}
	}

	// 处理 [0, 1, 2, 3] 这种情况，start 指向最后一个时，返回 start + 1
	if left == nums[left] {
		return left + 1
	}
	return left
}
