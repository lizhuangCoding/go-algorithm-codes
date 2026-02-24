package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/description/

func isCovered(ranges [][]int, left int, right int) bool {
	maxLen := 0
	for _, v := range ranges {
		if maxLen < v[1] {
			maxLen = v[1]
		}
	}

	diff := make([]int, maxLen+10)
	for i := 0; i < len(ranges); i++ {
		diff[ranges[i][0]]++
		diff[ranges[i][1]+1]--
	}

	target := make([]int, maxLen+10)
	for i := 1; i < len(diff); i++ {
		target[i] += target[i-1] + diff[i]
	}

	for i := left; i <= right; i++ {
		if i < 0 || i >= len(target) || target[i] == 0 {
			return false
		}
	}
	return true
}

// 代码简介度优化：
// func isCovered(ranges [][]int, left, right int) bool {
// 	diff := [52]int{} // 差分数组
// 	for _, r := range ranges {
// 		diff[r[0]]++
// 		diff[r[1]+1]--
// 	}
// 	cnt := 0 // 前缀和
// 	for i := 1; i <= 50; i++ {
// 		cnt += diff[i]
// 		if cnt <= 0 && left <= i && i <= right {
// 			return false
// 		}
// 	}
// 	return true
// }
