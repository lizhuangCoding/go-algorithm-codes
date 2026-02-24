package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/k-radius-subarray-averages/description/

func getAverages(nums []int, k int) []int {
	var arr = make([]int, 0)
	if len(nums) < 2*k+1 {
		for i := 0; i < len(nums); i++ {
			arr = append(arr, -1)
		}
		return arr
	}

	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		// 开始
		if right < k {
			arr = append(arr, -1)
		}
		// 中间
		if right >= 2*k {
			arr = append(arr, sum/(2*k+1))
			left := right - 2*k
			sum -= nums[left]
		}
	}
	// 结尾
	for right := len(nums) - k; right < len(nums); right++ {
		arr = append(arr, -1)
	}

	return arr
}
