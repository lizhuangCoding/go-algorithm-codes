package main

import (
	"math"
)

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/minimum-size-subarray-in-infinite-array/description/

func minSizeSubarray(nums []int, target int) int {
	l := len(nums)
	total := 0
	for i := 0; i < l; i++ {
		total += nums[i]
	}

	minLen := math.MaxInt64
	left := 0
	sum := 0
	for right := 0; right < l*2; right++ {
		sum += nums[right%l]
		for sum > target%total {
			sum -= nums[left%l]
			left++
		}

		if sum == target%total {
			if minLen > right-left+1 {
				minLen = right - left + 1
			}
		}
	}
	if minLen == math.MaxInt64 {
		return -1
	}
	return minLen + target/total*l
}

// func main() {
// 	fmt.Println(minSizeSubarray([]int{1, 1, 1, 2, 3}, 4))
// 	fmt.Println(minSizeSubarray([]int{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1}, 83))
// }
