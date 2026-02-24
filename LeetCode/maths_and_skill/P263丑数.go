package main

// 签到：数学
// 力扣：https://leetcode.cn/problems/ugly-number/

// 时间复杂度：O(logn)。时间复杂度取决于对 n 除以 2,3,5 的次数，由于每次至少将 n 除以 2，因此除法运算的次数不会超过 O(logn)。
// 空间复杂度：O(1)。
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
