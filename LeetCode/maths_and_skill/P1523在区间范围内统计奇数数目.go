package main

// 数学
// 力扣：https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/description/

// 方法一：暴力枚举
// func countOdds(low int, high int) int {
// 	num := 0
// 	for i := low; i <= high; i++ {
// 		if i%2 == 1 {
// 			num++
// 		}
// 	}
// 	return num
// }

// 方法二：找规律
func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 { // 两边都是偶数
		return (high - low) / 2
	} else if low%2 == 0 && high%2 == 1 { // 左边偶数，右边奇数
		return (high - low + 1) / 2
	} else if low%2 == 1 && high%2 == 0 { // 左边奇数，右边偶数
		return (high - low + 1) / 2
	} else { // 两边都是奇数
		return (high-low)/2 + 1
	}
}

// 方法三：前缀和
// pre(x) 为区间 [0,x] 中奇数的个数，所以 pre(x)=(x+1)/2，所以答案为 pre(high)−pre(low−1)。
// func countOdds(low int, high int) int {
// 	return preCountOdds(high) - preCountOdds(low-1)
// }
//
// // 找到[0, a]中奇数的个数
// func preCountOdds(a int) int {
// 	return (a + 1) >> 1
// }
