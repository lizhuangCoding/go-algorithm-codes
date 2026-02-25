package main

// 二分查找
// 力扣：https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/

func countNegatives(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		arr := grid[i]

		left, right := 0, len(arr)-1
		index := -1
		for left <= right {
			mid := (left + right) >> 1
			if arr[mid] < 0 {
				index = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if index != -1 {
			res += len(arr) - index
		}
	}

	return res
}
