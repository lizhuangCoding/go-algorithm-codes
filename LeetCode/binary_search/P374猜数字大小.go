package main

// 二分查找
// 力扣：https://leetcode.cn/problems/guess-number-higher-or-lower/

// func guessNumber(n int) int {
// 	left, right := 1, n
//
// 	for left <= right {
// 		mid := (left + right) / 2
//
// 		demo := guess(mid)
// 		if demo == 0 {
// 			return mid
// 		} else if demo == -1 {
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}
//
// 	return -1
// }
