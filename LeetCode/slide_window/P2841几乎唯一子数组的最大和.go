package main

// 滑动窗口，map
// 力扣：https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/

func maxSum(nums []int, m int, k int) int64 {
	if m > k {
		return 0
	}

	mm := make(map[int]int) // 判断是否符合m个数值
	result := 0             // 返回结果
	demo := 0               // 临时值存储滑动窗口的和
	for right := 0; right < len(nums); right++ {
		demo += nums[right]
		mm[nums[right]]++

		if right >= k-1 {
			// 需要符合m条件，而且值大于result
			if len(mm) >= m && result < demo {
				result = demo
			}

			left := right - (k - 1)
			demo -= nums[left]
			mm[nums[left]]--
			if mm[nums[left]] == 0 {
				delete(mm, nums[left])
			}
		}
	}
	return int64(result)
}
