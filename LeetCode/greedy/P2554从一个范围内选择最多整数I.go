package main

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-number-of-integers-to-choose-from-a-range-i/description/

func maxCount(banned []int, n int, maxSum int) int {
	m := make(map[int]bool)
	for _, v := range banned {
		m[v] = true
	}

	res := 0
	demo := 0
	for i := 1; i <= n && demo <= maxSum; i++ {
		if !m[i] && demo+i <= maxSum {
			demo += i
			res++
		}
	}

	return res
}
