package main

// 位运算
// 力扣：https://leetcode.cn/problems/insert-into-bits-lcci/
// b站视频：https://www.bilibili.com/video/BV1YT4y117AH?spm_id_from=333.788.videopod.episodes&vd_source=e2f6edb68f421b13075178535968e307&p=11

func insertBits(N, M, i, j int) int {
	// 先将第i到j位变为0，然后替换第i到j位
	for k := i; k <= j; k++ {
		// 在 ^(1 << k) 中，^ 是按位取反运算符，它的作用是对操作数的二进制位进行取反操作，即把二进制位中的 0 变为 1，1 变为 0 。
		N &= ^(1 << k)
	}
	return N | (M << i)
}
