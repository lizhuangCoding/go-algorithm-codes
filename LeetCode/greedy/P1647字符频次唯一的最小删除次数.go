package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-deletions-to-make-character-frequencies-unique/

func minDeletions(s string) int {
	// 记录字符次数
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a'] += 1
	}

	// 从大到小排序（字符次数）
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})

	res := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == 0 {

			res += arr[i] // 前一个是0，那么后面的需要全部删除（后面的需要比前面的更小）
			arr[i] = 0

		} else {

			// if arr[i] == arr[i-1] {
			// 	arr[i]--
			// 	res++
			// } else if arr[i] > arr[i-1] {
			// 	res += arr[i] - arr[i-1] + 1
			// 	arr[i] = arr[i-1] - 1
			// }
			// 可以合并为：
			if arr[i] >= arr[i-1] {
				res += arr[i] - arr[i-1] + 1
				arr[i] = arr[i-1] - 1
			}
		}
	}
	return res
}
