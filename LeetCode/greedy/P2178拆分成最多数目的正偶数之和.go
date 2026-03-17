package main

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-split-of-positive-even-integers/

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}

	res := make([]int64, 0)
	var i int64 = 1

	for finalSum != 0 {
		demo := i * 2
		if finalSum-demo <= demo {
			demo = finalSum
		}
		finalSum -= demo

		res = append(res, demo)
		i++
	}

	return res
}
