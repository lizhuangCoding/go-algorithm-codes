package main

// 二分查找
// 力扣：https://leetcode.cn/problems/arranging-coins/description/

// 方法一：暴力
// func arrangeCoins(n int) int {
// 	res := 1
// 	demo := 0
// 	for demo < n {
// 		demo += 1
// 		res++
// 	}
// 	return res - 1
// }

// 有点不理解:
// 方法二：二分查找
// func arrangeCoins(n int) int {
// 	return sort.Search(n, func(k int) bool {
// 		k++
// 		return k*(k+1) > 2*n
// 	})
// }
