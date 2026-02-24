package main

import (
	"strconv"
)

// 模拟
// 力扣：https://leetcode.cn/problems/multiply-strings/description/
// 题解：https://leetcode.cn/problems/multiply-strings/solutions/372098/zi-fu-chuan-xiang-cheng-by-leetcode-solution/

func multiply(s1 string, s2 string) string {
	if s1 == "0" || s2 == "0" {
		return "0"
	}

	result := "0"
	m, n := len(s1), len(s2)

	for i := n - 1; i >= 0; i-- {
		curr := ""

		// 先补0
		for j := n - 1; j > i; j-- {
			curr += "0"
		}

		y := int(s2[i] - '0')
		carry := 0 // 是否有进位

		for j := m - 1; j >= 0; j-- {
			x := int(s1[j] - '0')

			demo := x*y + carry
			carry = demo / 10
			curr = strconv.Itoa(demo%10) + curr
		}

		// 还有一个进位
		if carry != 0 {
			curr = strconv.Itoa(carry) + curr
		}

		// 两个字符串相加
		result = addStringMultiply(result, curr)
	}
	return result
}

// 两个字符串相加
func addStringMultiply(s1 string, s2 string) string {
	if s1 == "0" || s2 == "0" {
		return "0"
	}
	res := ""
	index1, index2 := len(s1)-1, len(s2)-1
	carry := 0

	for index1 >= 0 || index2 >= 0 {
		num1, num2 := 0, 0

		if index1 >= 0 {
			num1 = int(s1[index1] - '0')
			index1--
		}
		if index2 >= 0 {
			num2 = int(s2[index2] - '0')
			index2--
		}

		sum := num1 + num2 + carry
		sum, carry = sum%10, sum/10

		res = strconv.Itoa(sum) + res
	}

	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

// 或者：
// func multiply(num1 string, num2 string) string {
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	m, n := len(num1), len(num2)
// 	// ansArr 用来存储中间计算结果，每一位表示最终乘积的某一位数字
// 	ansArr := make([]int, m+n) // 两个整数相乘时，结果的位数最多为 m + n 位
//
// 	// 模拟竖式乘法
// 	for i := m - 1; i >= 0; i-- {
// 		x := int(num1[i]) - '0'
// 		for j := n - 1; j >= 0; j-- {
// 			y := int(num2[j] - '0')
// 			// 计算 num1[i] 和 num2[j] 的乘积，并将结果累加到 ansArr[i + j + 1]
// 			ansArr[i+j+1] += x * y
// 		}
// 	}
//
// 	// 处理进位
// 	for i := m + n - 1; i > 0; i-- {
// 		ansArr[i-1] += ansArr[i] / 10
// 		ansArr[i] %= 10
// 	}
//
// 	// 构造结果字符串
// 	result := ""
// 	idx := 0
// 	// 如果最高位 ansArr[0] 为 0，则跳过
// 	if ansArr[0] == 0 {
// 		idx = 1
// 	}
// 	// 遍历 ansArr，将每一位数字转换为字符串并拼接到结果字符串 ans
// 	for ; idx < m+n; idx++ {
// 		result += strconv.Itoa(ansArr[idx])
// 	}
// 	return result
// }
