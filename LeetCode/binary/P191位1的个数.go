package main

// 二进制，求二进制中1的个数
// 力扣：https://leetcode.cn/problems/number-of-1-bits/

// 方法一：右移输入的数字，判断每一次的末尾是否是1
// 因为n是正数，所以可以像下面这样写。如果n为负数，还按照下面的写法就会陷入死循环（二进制的第一位是1，负数右移会一直存在1），这种情况怎么搞？
// func hammingWeight(n int) int {
// 	count := 0
//
// 	for n != 0 {
// 		if n&1 == 1 {
// 			count++
// 		}
// 		n >>= 1
// 	}
//
// 	return count
// }

// 方法二：设置一个中间值1，让这个值一直左移
// func hammingWeight(n int) int {
// 	count := 0
// 	flag := 1
//
// 	// 当flag超出int的范围后就是负数了，这个时候就可以退出了
// 	for flag > 0 {
// 		if n&flag > 0 {
// 			count++
// 		}
// 		flag <<= 1
// 	}
//
// 	// 当然，也可以手动的设置循环次数
// 	// for i := 0; i < 32; i++ {
// 	// 	if n&flag > 0 {
// 	// 		count++
// 	// 	}
// 	// 	flag <<=   1
// 	// }
// 	return count
// }

// 方法三：每一次循环都是消除n的二进制中的最后一个1。
// 二进制中有几个1就遍历几次。
func hammingWeight(n int) int {
	count := 0

	for n != 0 {
		count++
		n = n & (n - 1) // 用&来消除n的二进制中的最后一个1
	}

	return count
}
