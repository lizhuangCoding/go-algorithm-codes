package main

// 位运算
// 力扣：https://leetcode.cn/problems/binary-number-with-alternating-bits/

// 方法一：模拟
// func hasAlternatingBits(n int) bool {
// 	var flag bool // true表示当前位置应该是1，false表示当前位置应该是2，如果不满足return false
// 	// 初始化flag
// 	if n&1 == 1 {
// 		flag = true
// 	} else {
// 		flag = false
// 	}
//
// 	// 开始判断
// 	for n != 0 {
// 		if !((flag && n&1 == 1) || (!flag && n&1 == 0)) {
// 			return false // 如果条件不满足就return
// 		}
// 		n >>= 1
// 		flag = !flag
// 	}
// 	return true
// }

// 类似于上面的写法：模拟
// 时间复杂度：O(logn)。输入 n 的二进制表示最多有 O(logn) 位。
// 空间复杂度：O(1)。使用了常数空间来存储中间变量。

// func hasAlternatingBits(n int) bool {
// 	pre := 2
// 	for ; n != 0; n /= 2 {
// 		cur := n % 2
// 		if cur == pre {
// 			return false
// 		}
// 		pre = cur
// 	}
// 	return true
// }

// 方法二：位运算，异或
// 时间复杂度：O(1)。空间复杂度：O(1)
func hasAlternatingBits(n int) bool {
	a := n ^ n>>1 // 如果n是正确的二进制，那么a的二进制应该全是1
	return a&(a+1) == 0
}
