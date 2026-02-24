package main

// 滑动窗口，map
// 力扣：https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/

func maximumSubarraySum(nums []int, k int) int64 {
	m := make(map[int]int)

	maxResult := 0
	demo := 0
	for right := 0; right < len(nums); right++ {
		demo += nums[right]
		m[nums[right]]++

		if right >= k-1 {
			if len(m) == k && demo > maxResult {
				maxResult = demo
			}

			left := right - (k - 1)
			demo -= nums[left]
			m[nums[left]]--
			if m[nums[left]] == 0 {
				delete(m, nums[left])
			}
		}
	}
	return int64(maxResult)
}
