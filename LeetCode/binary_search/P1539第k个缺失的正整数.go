package main

// 二分查找
// 力扣：https://leetcode.cn/problems/kth-missing-positive-number/

// 方法一：枚举
// func findKthPositive(arr []int, k int) int {
// 	res := 0
// 	demo := 1
//
// 	for i := 0; i < len(arr); {
//
// 		if arr[i] != demo {
// 			k--
// 			if k == 0 {
// 				res = demo
// 				break
// 			}
// 		} else {
// 			i++
// 		}
// 		demo++
// 	}
//
// 	if k != 0 {
// 		res = k + arr[len(arr)-1]
// 	}
//
// 	return res
// }

// 方法二：二分查找
// 思路：
// 数组：1, 2, 3, 4, 5
// 下标：0, 1, 2, 3, 4
// 不缺失正整数的情况下：arr[i] - i - 1 == 0 是成立的
// 数组：2, 3, 5, 7, 11
// 下标：0, 1, 2, 3, 4
// 缺失正整数的情况下：arr[i] - i - 1 得到的数字，就是 arr[i] 之前缺失的正整数的数量
// 因此我们可以采用二分的方式，找到最左边的那个满足表达式 arr[i] - i - 1 >= k 的下标
// 用得到的下标 + k 即可得到缺失的第 k 个正整数
func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1

	index := len(arr)
	for left <= right {
		mid := (left + right) >> 1

		if arr[mid]-mid-1 >= k { // 缺失的数多，往左边走
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index + k
}

// 或者：
// func findKthPositive(arr []int, k int) int {
// 	left, right := 0, len(arr)
//
// 	// 二分查找：找到第一个满足"缺失数数量 >= k"的位置
// 	for left < right {
// 		mid := left + (right-left)/2
//
// 		// 计算到当前位置为止缺失的正整数数量
// 		// arr[mid]是当前位置的值，下标从0开始
// 		// 正常情况下，如果没有任何缺失，arr[mid]应该是mid+1
// 		// 所以缺失数量 = arr[mid] - (mid+1) = arr[mid] - mid - 1
// 		missingCount := arr[mid] - mid - 1
//
// 		if missingCount >= k {
// 			right = mid // 缺失数足够多，向左搜索
// 		} else {
// 			left = mid + 1 // 缺失数不够，向右搜索
// 		}
// 	}
//
// 	// 最终left指向第一个缺失数数量>=k的位置
// 	// 第k个缺失的数 = left + k
// 	return left + k
// }
