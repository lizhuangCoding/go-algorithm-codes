package main

// 位运算
// 力扣：https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c/description/

// 模拟：
// func minFlips(a int, b int, c int) int {
// 	res := 0
//
// 	for a != 0 || b != 0 || c != 0 {
// 		a1, b1, c1 := 0, 0, 0
// 		if a&1 == 1 {
// 			a1 = 1
// 		}
//
// 		if b&1 == 1 {
// 			b1 = 1
// 		}
//
// 		if c&1 == 1 {
// 			c1 = 1
// 		}
//
// 		if a1|b1 != c1 {
// 			if a1 == 1 && b1 == 1 { // 这种情况两个都需要改变（c1 == 0）
// 				res += 2
// 			} else {
// 				res++
// 			}
// 		}
//
// 		a >>= 1
// 		b >>= 1
// 		c >>= 1
// 	}
// 	return res
// }

// 优化：
func minFlips(a int, b int, c int) int {
	res := 0

	for i := 0; i < 31; i++ {
		a1 := (a >> i) & 1
		b1 := (b >> i) & 1
		c1 := (c >> i) & 1

		if c1 == 0 {
			res += a1 + b1 // 如果a1是1，b1也是1，那么就加2。
		} else {
			if a1 == 0 && b1 == 0 {
				res++
			}
		}

	}
	return res
}
