package main

import "math"

// 数组
// 力扣：https://leetcode.cn/problems/find-closest-lcci/description/

func findClosest(words []string, word1 string, word2 string) int {
	minRes := math.MaxInt32

	ind1, ind2 := -1, -1

	for i := 0; i < len(words); i++ {

		if words[i] == word1 {
			ind1 = i
		} else if words[i] == word2 {
			ind2 = i
		}

		if ind1 != -1 && ind2 != -1 {
			minRes = int(math.Min(float64(minRes), math.Abs(float64(ind2-ind1))))
		}
	}
	return minRes
}
