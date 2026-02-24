package main

// 数学：脑筋急转弯
// 力扣：https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements/description/

func maximizeSum(nums []int, k int) int {
	maxVal := 0
	for i := 0; i < len(nums); i++ {
		if maxVal < nums[i] {
			maxVal = nums[i]
		}
	}
	return myMaximizeSum(k-1) + k*maxVal
}

// 求 n -> 1 + 2 + 3 + ... + n
func myMaximizeSum(n int) int {
	// res := 0
	// for i := 1; i <= n; i++ {
	// 	res += i
	// }
	// return res

	// 优化：找规律
	if n%2 == 0 {
		return (n + 1) * (n / 2)
	} else {
		return (n+1)/2 + (n+1)*(n/2)
	}
}

// 或者：
// func maximizeSum(nums []int, k int) int {
// 	return (slices.Max(nums)*2 + k - 1) * k / 2
// }
