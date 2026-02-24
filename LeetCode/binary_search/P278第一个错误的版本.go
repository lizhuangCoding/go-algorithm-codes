package main

// 二分查找：最左边
// 力扣：https://leetcode.cn/problems/first-bad-version/

// func firstBadVersion(n int) int {
// 	left, right := 1, n
// 	result := 0
//
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if isBadVersion(mid) {
// 			result = mid
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return result
// }

// // 或者：
// func firstBadVersion(n int) int {
// 	// sort.Search 是 Go 标准库中 sort 包提供的一个通用二分查找函数
// 	return sort.Search(n, func(version int) bool {
// 		return isBadVersion(version)
// 	})
// }
