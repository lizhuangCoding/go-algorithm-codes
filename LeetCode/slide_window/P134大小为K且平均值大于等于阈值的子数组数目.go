package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
// func numOfSubarrays(arr []int, k int, threshold int) int {
// 	num := 0
//
// 	sum := 0
// 	for i := 0; i <= k-1; i++ {
// 		sum += arr[i]
// 	}
// 	if sum/k >= threshold {
// 		num++
// 	}
//
// 	for right := k; right < len(arr); right++ {
// 		left := right - k
// 		sum -= arr[left]
// 		sum += arr[right]
// 		if sum/k >= threshold {
// 			num++
// 		}
// 	}
// 	return num
// }

// 或者：
func numOfSubarrays(arr []int, k int, threshold int) int {
	num := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		if i < k-1 {
			sum += arr[i]
		} else {
			sum += arr[i]
			if sum/k >= threshold {
				num++
			}

			left := i - k + 1
			sum -= arr[left]
		}
	}
	return num
}
