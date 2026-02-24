package main

// 分组循环，数组
// 力扣：https://leetcode.cn/problems/longest-alternating-subarray/

func alternatingSubarray(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		start := i
		demo := 1 // 一开始

		for ; i < len(nums)-1 && nums[i]+demo == nums[i+1]; i++ {
			demo *= -1
		}

		result = alternatingSubarrayMax(result, i-start+1)

		// 回退一次
		// 例子：2, 3, 4, 3, 4。是为了避免走完 2，3 后直接开始走 4，3，让i--回退一次，下一次循环会从3开始，而不是从4开始
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			i--
		}

	}

	if result <= 1 {
		return -1
	}
	return result
}

func alternatingSubarrayMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(alternatingSubarray([]int{2, 3, 4, 3, 4})) // 4
// }
