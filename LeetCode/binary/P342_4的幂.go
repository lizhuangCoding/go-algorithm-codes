package main

// 数学

// 取模，首先判断n是否是2的幂（n&(n-1) == 0），在此基础上再判断n是否是4的幂。
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
