package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/maximum-population-year/description/

func maximumPopulation(logs [][]int) int {
	diff := make([]int, 2060)

	for i := 0; i < len(logs); i++ {
		diff[logs[i][0]]++
		// 第 logs[i][1] 年，这个人已经去世了，不能算上这个人还活着，所以不用 diff[logs[i][1]+1]--，要从第logs[i][1]年就开始--
		diff[logs[i][1]]--
	}

	result := 0
	maxSum := 0
	cnt := 0
	for i := 1950; i <= 2050; i++ {
		cnt += diff[i]
		if maxSum < cnt {
			maxSum = cnt
			result = i
		}
	}
	return result
}
