package main

// 位运算
// 力扣：https://leetcode.cn/problems/check-if-bitwise-or-has-trailing-zeros/description/

// 找到两个偶数就可以了
func hasTrailingZeros(nums []int) bool {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 != 1 { // 偶数
			n++
		}
		if n >= 2 {
			return true
		}
	}

	return false
}
