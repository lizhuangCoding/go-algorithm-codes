package main

import (
	"math"
)

// 方法一：暴力，时间超限
// func nonSpecialCount(l int, r int) int {
// 	num := 0
// 	for i := l; i <= r; i++ {
// 		if judgeNonSpecialCount(i) {
// 			num++
// 		}
// 	}
// 	return r - l + 1 - num
// }
//
// func judgeNonSpecialCount(a int) bool {
// 	n := 0
// 	for i := 1; i < a; i++ {
// 		if a%i == 0 {
// 			n++
// 		}
// 		if n > 2 {
// 			return false
// 		}
// 	}
// 	return n == 2
// }

// 方法二：质数筛
func nonSpecialCount(l int, r int) int {
	n := int(math.Sqrt(float64(r)))
	v := make([]int, n+1)

	res := r - l + 1
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			if i*i >= l && i*i <= r { // 特殊数字是i*i
				res--
			}
			for j := i * 2; j <= n; j += i { // 非特殊数字
				v[j] = 1
			}
		}
	}

	return res
}
