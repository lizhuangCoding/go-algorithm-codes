package main

// 二进制
// 力扣：https://leetcode.cn/problems/sum-of-values-at-indices-with-k-set-bits/

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if countSumIndicesWithKSetBits(i) == k {
			res += nums[i]
		}
	}
	return res
}

// 计算数字n的二进制中有几个1
func countSumIndicesWithKSetBits(n int) int {
	res := 0
	for n != 0 {
		if n%2 == 1 {
			res++
		}
		n /= 2
	}
	return res
}
