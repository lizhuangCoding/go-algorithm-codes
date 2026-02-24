package main

// 位运算
// 力扣：https://leetcode.cn/problems/find-the-k-or-of-an-array/description/

func findKOr(nums []int, k int) int {
	maxNum := 0 // 最大值
	for i := 0; i < len(nums); i++ {
		maxNum = max(maxNum, nums[i])
	}

	res := 0
	demo := 1

	// 如果枚举的值大于数组中的最大值就退出
	for maxNum >= demo {
		numDemo := 0 // 目前符合要求的数量

		for i := 0; i < len(nums); i++ {

			if nums[i] >= demo {
				if nums[i]&demo == demo { // 符合要求的元素
					numDemo++
				}
			}

			if numDemo == k {
				res = res | demo
				break
			}
		}
		demo <<= 1
	}

	return res
}

// // 或者：
// func findKOr(nums []int, k int) int {
// 	ans := 0
//
// 	for i := 0; i < 31; i++ {
// 		cnt := 0
//
// 		for _, num := range nums {
// 			if (num>>i)&1 == 1 {
// 				cnt++
// 			}
// 		}
//
// 		if cnt >= k {
// 			ans |= 1 << i
// 		}
// 	}
// 	return ans
// }
