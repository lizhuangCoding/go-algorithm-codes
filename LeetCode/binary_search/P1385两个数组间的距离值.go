package main

import "slices"

// 二分查找
// 力扣：https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/

// 思路：只要找到大于等于 arr1[i] 的第一个 arr2[j] 和 小于 arr1[i] 的第一个 arr2[j]。
// 判断这两个值是否满足就可以了，如果这两个值都满足，那么其他的值一定满足

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)

	res := 0
	for i := 0; i < len(arr1); i++ {
		index := xiaoyuFindTheDistanceValue(arr2, arr1[i])

		if index != -1 && index+1 <= len(arr2)-1 {
			if absFindTheDistanceValue(arr1[i]-arr2[index]) > d && absFindTheDistanceValue(arr1[i]-arr2[index+1]) > d {
				res++
			}
		} else if index != -1 {
			if absFindTheDistanceValue(arr1[i]-arr2[index]) > d {
				res++
			}
		} else if index == -1 {
			if absFindTheDistanceValue(arr1[i]-arr2[index+1]) > d {
				res++
			}
		}
	}

	return res
}

// 小于 arr1[i] 的第一个 arr2[j]。
func xiaoyuFindTheDistanceValue(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			left = mid + 1
		}
	}
	return index
}

func absFindTheDistanceValue(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
