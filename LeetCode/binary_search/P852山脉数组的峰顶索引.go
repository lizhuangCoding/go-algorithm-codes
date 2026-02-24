package main

// 二分查找
// 力扣：https://leetcode.cn/problems/peak-index-in-a-mountain-array/description/
// 和题目162相同

// 时间复杂度O(n)
// func peakIndexInMountainArray(arr []int) int {
// 	for i := 1; ; i++ {
// 		if arr[i] > arr[i+1] {
// 			return i
// 		}
// 	}
// }

// 方法二：二分。时间复杂度O(logn)
func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2
		// 刚好在顶峰
		if (mid-1 >= 0 && arr[mid] > arr[mid-1]) && (mid+1 < n && arr[mid] > arr[mid+1]) {
			return mid
		} else if (mid-1 >= 0 && arr[mid] > arr[mid-1]) && (mid+1 < n && arr[mid] < arr[mid+1]) { // 在顶峰的左边
			left = mid
		} else if (mid-1 >= 0 && arr[mid] < arr[mid-1]) && (mid+1 < n && arr[mid] > arr[mid+1]) { // 在顶峰的右边
			right = mid
		}
	}
	return -1
}

// 或者：
// func peakIndexInMountainArray(arr []int) int {
// 	l, r := 0, len(arr)-1
//
// 	for l <= r {
// 		mid := (l + r) >> 1
// 		if mid-1 >= 0 && mid+1 <= len(arr)-1 && arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
// 			return mid
// 		} else if (mid-1 >= 0 && arr[mid] < arr[mid-1]) || (mid+1 <= len(arr)-1 && arr[mid] > arr[mid+1]) {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
//
// 	return -1
// }

// 或者：
// func peakIndexInMountainArray(arr []int) int {
// 	left, right := 0, len(arr)-2
// 	for left+1 < right {
// 		mid := left + (right-left)/2
// 		if arr[mid] > arr[mid+1] {
// 			right = mid
// 		} else {
// 			left = mid
// 		}
// 	}
// 	return right
// }

// // 或者：使用库函数
// func peakIndexInMountainArray(arr []int) int {
// 	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
// }
