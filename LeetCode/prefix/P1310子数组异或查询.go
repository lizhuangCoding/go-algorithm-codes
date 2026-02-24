package main

// 前缀和、异或
// 力扣：https://leetcode.cn/problems/xor-queries-of-a-subarray/description/

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]

	res := make([]int, 0)

	for i := 1; i < len(arr); i++ {
		prefix[i] = arr[i] ^ prefix[i-1]
	}

	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]

		demo := 0
		if l > 0 {
			demo = prefix[r] ^ prefix[l-1]
		} else {
			demo = prefix[r]
		}
		res = append(res, demo)
	}

	return res
}
