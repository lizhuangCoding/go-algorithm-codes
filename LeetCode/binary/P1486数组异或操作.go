package main

// 签到：位运算
func xorOperation(n int, start int) int {
	demo := 0
	res := 0
	for i := 0; i < n; i++ {
		demo = start + 2*i
		res ^= demo
	}
	return res
}
