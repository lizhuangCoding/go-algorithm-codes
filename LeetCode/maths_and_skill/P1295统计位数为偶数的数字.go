package main

import "math"

// 数学
// 力扣：https://leetcode.cn/problems/find-numbers-with-even-number-of-digits/description/

// 方法一：
// func findNumbers(nums []int) int {
// 	res := 0
//
// 	for i := 0; i < len(nums); i++ {
// 		demo := 0
// 		for nums[i] > 0 {
// 			nums[i] /= 10
// 			demo++
// 		}
// 		if demo%2 == 0 {
// 			res++
// 		}
// 	}
// 	return res
// }

// 方法二：将数字转换为字符串，然后看字符串的长度是否为偶数
// 省略...

// 方法三：枚举 + 数学
func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		// math.Log10(float64(num)) 计算 num 的常用对数（底数为 10），并将数字转换为 float64 类型。
		// 对数结果实际上是 num 在 10 的底下的指数，也就是说，如果 num 是 100，那么 math.Log10(100) 将返回 2
		if int(math.Log10(float64(num)))+1%2 == 0 {
			ans++
		}
	}
	return ans
}
