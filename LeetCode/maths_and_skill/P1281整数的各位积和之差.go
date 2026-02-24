package main

import (
	"strconv"
)

// 签到：数学
func subtractProductAndSum(n int) int {
	num1, num2 := 1, 0
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		demo := int(str[i] - 49 + 1)
		num1 *= demo
		num2 += demo
	}
	return num1 - num2
}

// 或者：
// func subtractProductAndSum(n int) int {
// 	m := 1
// 	s := 0
// 	for n > 0 {
// 		x := n % 10
// 		n /= 10
// 		m *= x
// 		s += x
// 	}
// 	return m - s
// }
