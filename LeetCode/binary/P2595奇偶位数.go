package main

// 位运算
// 力扣：https://leetcode.cn/problems/number-of-even-and-odd-bits/description/

// 时间复杂度：O(logn)
func evenOddBit(n int) []int {
	even, odd := 0, 0

	isOushu := true

	for n > 0 {
		if isOushu && n&1 == 1 {
			even++
		}
		if !isOushu && n&1 == 1 {
			odd++
		}

		isOushu = !isOushu
		n >>= 1
	}
	return []int{even, odd}
}
