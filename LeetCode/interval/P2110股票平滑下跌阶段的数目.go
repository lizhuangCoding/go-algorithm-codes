package main

// 分组循环，等差数列
// 力扣：https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/description/

func getDescentPeriods(prices []int) int64 {
	var result int64

	for i := 0; i < len(prices); i++ {
		start := i
		for ; i < len(prices)-1 && prices[i]-1 == prices[i+1]; i++ {
		}

		n := i - start + 1
		result += int64(n * (n + 1) / 2)

	}
	return result
}

// 为什么可以用等差数列？
// 看 P1759统计同质子字符串的数目 中的注释。
