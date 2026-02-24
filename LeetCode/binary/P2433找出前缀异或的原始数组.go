package main

// 位运算（异或）
// 力扣：https://leetcode.cn/problems/find-the-original-array-of-prefix-xor/description/

func findArray(pref []int) []int {
	if len(pref) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	demo := 0
	for i := 0; i < len(pref); i++ {
		res = append(res, demo^pref[i])
		demo = demo ^ res[i]
	}
	return res
}
