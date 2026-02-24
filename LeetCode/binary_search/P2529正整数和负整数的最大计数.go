package main

// 二分查找
// 力扣：https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/
//
// func maximumCount(nums []int) int {
// 	i1, i2 := fuMaximumCount(nums), zhengMaximumCount(nums)
//
// 	n1 := i1 + 1
//
// 	n2 := -1
// 	if i2 != -1 {
// 		n2 = len(nums) - i2
// 	}
//
// 	if n1 > n2 {
// 		return n1
// 	}
//
// 	return n2
// }
//
// // 最右边的负数的索引位置
// func fuMaximumCount(nums []int) int {
// 	res := -1
// 	l, r := 0, len(nums)-1
//
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if nums[mid] >= 0 {
// 			r = mid - 1
// 		} else {
// 			res = mid
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }
//
// // 最左边的正数的索引位置
// func zhengMaximumCount(nums []int) int {
// 	res := -1
// 	l, r := 0, len(nums)-1
//
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if nums[mid] <= 0 {
// 			l = mid + 1
// 		} else {
// 			res = mid
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// 优化：
func maximumCount(nums []int) int {
	n1 := helperMaximumCount(nums, 0) // 找到小于0的值的最大索引位置
	n2 := helperMaximumCount(nums, 1) // 找到小于1的值的最大索引位置

	return max(n1+1, len(nums)-n2-1)
}

// 找到小于val的第一个索引位置
func helperMaximumCount(nums []int, target int) int {
	res := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			res = mid
			l = mid + 1
		}
	}
	return res
}
