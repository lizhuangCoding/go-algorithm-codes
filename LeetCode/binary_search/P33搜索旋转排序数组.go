package main

// 二分查找
// 力扣：https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] == target {
			return mid
		}

		// 如果从数组开头到中间的部分是有序的：[0, mid)
		if nums[0] <= nums[mid] {

			// 如果target在左边中，就更改right
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else { // 如果不在左边中，就更改left
				left = mid + 1
			}

		} else { // 如果从中间到数组末尾的部分是有序的：(mid, nums.len-1]

			// 如果target在右边中，就更改left
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else { // 如果不在右边中，就更改right
				right = mid - 1
			}
		}
	}

	return -1
}
