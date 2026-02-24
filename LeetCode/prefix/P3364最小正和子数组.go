package main

import (
	"math"
)

// 前缀和
// 力扣：https://leetcode.cn/problems/minimum-positive-sum-subarray/description/

func minimumSumSubarray(nums []int, l int, r int) int {
	// 前缀和
	pre := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	// fmt.Println(pre)

	result := math.MaxInt64

	// 遍历每个可能的子数组长度
	for i := l; i <= r; i++ {
		d1, d2 := 0, i
		// fmt.Println("-------", d1, d2)

		// 确保d2不超出范围
		for d2 < len(pre) {
			// fmt.Println("=======",d1, d2)
			demo := pre[d2] - pre[d1]
			if demo > 0 && demo < result {
				result = demo
			}

			d1++
			d2++
		}
	}

	if result == math.MaxInt64 {
		return -1
	}
	return result
}

// func main() {
// 	fmt.Println(minimumSumSubarray([]int{3, -2, 1, 4}, 2, 3)) // 1
// 	fmt.Println(minimumSumSubarray([]int{1, 2, 3, 4}, 2, 4))  // 3
// 	fmt.Println(minimumSumSubarray([]int{7, 3}, 2, 2)) // 10
// }
