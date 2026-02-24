package main

// 二分查找
// 力扣：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
// 注意：题目说没有重复的元素，所以可以用二分法；但是如果有重复的元素的话，比如 1，1，1，0，1，1 和 1，1，0，1，1，1 这种情况就无法判断最小的0在左边还是右边，这个时候就要特殊判定了。

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left <= right {
		mid = (left + right) / 2

		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
		// 为什么是 right = mid，left = mid+1？
		// 一个是mid，一个是mid+1，可以画个图：3 4 5 1 2，试一试不同的情况下，left和right应该怎么走
	}
	return nums[mid]
}

// 或者：
// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1
//
// 	for left < right {
// 		mid := (left + right) / 2
//
// 		if nums[mid] < nums[right] {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return nums[left] // 为什么返回的是left，而不是mid？因为有 2，1 的情况。
// }
