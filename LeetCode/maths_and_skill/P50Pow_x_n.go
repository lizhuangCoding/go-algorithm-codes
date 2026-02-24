package main

// 数学，快速幂
// 力扣：https://leetcode.cn/problems/powx-n/description/

// 方法一：暴力，时间超限
// func myPow(x float64, n int) float64 {
// 	if x == 1 || x == 0 {
// 		return x
// 	}
//
// 	if n < 0 {
// 		x = 1.0 / x
// 		n = -n
// 	}
//
// 	res := 1.0
// 	for i := 0; i < n; i++ {
// 		res *= x
// 	}
// 	return float64(res)
// }

// 方法二：快速幂
// 时间复杂度：O(log∣n∣)。
// 空间复杂度：O(1)。
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	}

	res := 1.0
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	return res
}

// // 快速幂模版：
// func ksm(a, b int) int {
// 	res := 1
// 	for b > 0 {
// 		if b%2 == 1 {
// 			res *= a
// 		}
// 		a *= a
// 		b /= 2
// 	}
// 	return res
// }

// func main() {
// 	fmt.Println(myPow(2.00000, 10)) // 1024.00000
// 	fmt.Println(myPow(2.10000, 3))  // 9.26100
// 	fmt.Println(myPow(2.00000, -2)) // 0.25000
// }
