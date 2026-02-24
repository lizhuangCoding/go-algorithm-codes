package main

// 数学，思维
// 力扣：https://leetcode.cn/problems/Ju9Xwi/

// 方法一：数学
func leastMinutes(n int) int {
	demo := 1
	res := 0

	// 需要加倍几次
	for demo < n {
		demo *= 2
		res++
	}
	// 下载1次
	res++
	return res
}

// 方法二：二进制。根据 2的次方 和 二进制位数 之间的关系
// func leastMinutes(n int) int {
// 	// bits.Len：返回给定无符号整数的位数（以二进制形式表示时所需的位数）
// 	return bits.Len(uint(n-1)) + 1 // bits.Len(uint(n-1))：加倍几次。+1：下载1次
// }
