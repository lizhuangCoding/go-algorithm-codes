package main

// 数组、排序
// 力扣：https://leetcode.cn/problems/sort-an-array/description/

// 归并排序：
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	n := len(nums) / 2
	left, right := nums[:n], nums[n:]
	return mergeSort(sortArray(left), sortArray(right))
}

func mergeSort(left, right []int) []int {
	result := make([]int, 0)

	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	if l < len(left) {
		result = append(result, left[l:]...)
	}

	if r < len(right) {
		result = append(result, right[r:]...)
	}

	return result
}
