package main

// 位运算
// 力扣：https://leetcode.cn/problems/hamming-distance/description/

// 方法一：不断右移，并比较最后一位是否相同
// func hammingDistance(x int, y int) int {
// 	count := 0
//
// 	for !(x == 0 && y == 0) {
// 		// x&1 == y&1：相同
// 		if !(x&1 == y&1) {
// 			count++
// 		}
//
// 		x >>= 1
// 		y >>= 1
// 	}
//
// 	return count
// }

// 方法二：利用异或的特点
func hammingDistance(x int, y int) int {
	count := 0
	n := x ^ y

	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}

	return count
}
