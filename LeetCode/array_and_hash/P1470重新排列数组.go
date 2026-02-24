package main

// 签到：数组

func shuffle(nums []int, n int) []int {
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return res
}
