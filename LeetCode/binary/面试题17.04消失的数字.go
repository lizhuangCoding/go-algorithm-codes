package main

// 签到：异或
func missingNumber(nums []int) int {
	n := len(nums)

	res := 0
	for i := 0; i < n; i++ {
		res ^= nums[i]
		res ^= i
	}
	res ^= n

	return res
}
