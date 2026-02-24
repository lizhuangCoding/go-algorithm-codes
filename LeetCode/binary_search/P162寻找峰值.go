package main

// 二分查找
// 力扣：https://leetcode.cn/problems/find-peak-element/description/

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	mid := (l + r) >> 1

	for l <= r {
		mid = (l + r) >> 1

		if mid-1 != -1 && mid+1 <= len(nums)-1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if (mid-1 != -1 && nums[mid] > nums[mid-1]) || (mid+1 <= len(nums)-1 && nums[mid] < nums[mid+1]) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return mid
}
