package main

import "strconv"

// 十进制转换为七进制
// 力扣：https://leetcode.cn/problems/base-7/description/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""
	// 是否为正数
	isZheng := true
	if num < 0 {
		isZheng = false
		num = -num
	}

	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}

	if !isZheng {
		res = "-" + res
	}
	return res
}
