package main

// 双指针（二分查找）
// 力扣：https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

// func twoSum(numbers []int, target int) []int {
// 	if len(numbers) <= 1 {
// 		return nil
// 	}
//
// 	for index1 := 0; index1 < len(numbers); index1++ {
// 		for index2 := index1 + 1; index2 < len(numbers); index2++ {
// 			if numbers[index1]+numbers[index2] == target {
// 				return []int{index1 + 1, index2 + 1}
// 			} else if numbers[index1]+numbers[index2] > target {
// 				break
// 			}
// 		}
// 	}
//
// 	return nil
// }

// 优化：双指针
// 时间复杂度：O(n)，其中 n 是数组的长度。两个指针移动的总次数最多为 n 次。
// 空间复杂度：O(1)。
func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1

	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum > target {
			high--
		} else {
			low++
		}
	}

	return []int{-1, -1}
}

// 也可以二分查找：
// 时间复杂度：O(nlogn)，其中 n 是数组的长度。需要遍历数组一次确定第一个数，时间复杂度是 O(n)，寻找第二个数使用二分查找，时间复杂度是 O(logn)，因此总时间复杂度是 O(nlogn)。
// 空间复杂度：O(1)。
// func twoSum(numbers []int, target int) []int {
// 	for i := 0; i < len(numbers); i++ {
// 		low, high := i+1, len(numbers)-1
// 		for low <= high {
// 			mid := (high-low)/2 + low
// 			if numbers[mid] == target-numbers[i] {
// 				return []int{i + 1, mid + 1}
// 			} else if numbers[mid] > target-numbers[i] {
// 				high = mid - 1
// 			} else {
// 				low = mid + 1
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }
