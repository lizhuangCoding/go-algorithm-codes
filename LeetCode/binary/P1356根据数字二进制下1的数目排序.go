package main

import (
	"math/bits"
	"sort"
)

// 位运算
// 力扣：https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/description/

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if a != b {
			return a < b
		}
		return arr[i] < arr[j]
	})

	return arr
}

// // 或者：纯手搓
// func sortByBits(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}
//
// 	// 快速排序
// 	quickSortByBits(arr, 0, len(arr)-1)
// 	return arr
// }
//
// // 快速排序实现
// func quickSortByBits(arr []int, left, right int) {
// 	if left < right {
// 		pivotIndex := partitionByBits(arr, left, right)
// 		quickSortByBits(arr, left, pivotIndex-1)
// 		quickSortByBits(arr, pivotIndex+1, right)
// 	}
// }
//
// // 分区函数
// func partitionByBits(arr []int, left, right int) int {
// 	pivot := arr[right]
// 	i := left - 1
//
// 	for j := left; j < right; j++ {
// 		if compareByBits(arr[j], pivot) {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}
//
// 	arr[i+1], arr[right] = arr[right], arr[i+1]
// 	return i + 1
// }
//
// // 比较函数：先比较二进制中1的个数，再比较数值本身
// func compareByBits(a, b int) bool {
// 	countA := countOnes(a)
// 	countB := countOnes(b)
//
// 	if countA != countB {
// 		return countA < countB
// 	}
// 	return a < b
// }
//
// // 计算数字二进制中1的个数
// func countOnes(num int) int {
// 	count := 0
// 	n := num
//
// 	for n != 0 {
// 		count += n & 1 // 检查最低位是否为1
// 		n >>= 1        // 右移一位
// 	}
//
// 	return count
// }
