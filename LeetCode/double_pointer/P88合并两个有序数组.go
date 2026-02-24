package main

// 双指针
// LeetCode：https://leetcode.cn/problems/merge-sorted-array/

// 方法一：直接合并后排序
// 时间复杂度：O((m+n)log(m+n))。排序序列长度为 m+n，套用快速排序的时间复杂度即可，平均情况为 O((m+n)log(m+n))。
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	copy(nums1[m:], nums2)
// 	sort.Ints(nums1)
// 	// 或者：
// 	// sort.Slice(nums1, func(i, j int) bool {
// 	// 	return nums1[i] < nums1[j]
// 	// })
// }

// 方法二：double_pointer
func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, 0, m+n)
	ind1, ind2 := 0, 0

	for {
		if ind1 == m {
			arr = append(arr, nums2[ind2:]...)
			break
		}
		if ind2 == n {
			arr = append(arr, nums1[ind1:]...)
			break
		}

		if nums1[ind1] <= nums2[ind2] {
			arr = append(arr, nums1[ind1])
			ind1++
		} else {
			arr = append(arr, nums2[ind2])
			ind2++
		}
	}
	copy(nums1, arr)
}
