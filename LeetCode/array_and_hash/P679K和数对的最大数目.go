package main

// 哈希
// 力扣：https://leetcode.cn/problems/max-number-of-k-sum-pairs/description/

func maxOperations(nums []int, k int) int {
	m := make(map[int]int)

	res := 0
	for i := 0; i < len(nums); i++ {
		demo := nums[i]
		if m[k-demo] > 0 {
			m[k-demo]--
			res++
		} else {
			m[demo]++
		}
	}

	return res
}
