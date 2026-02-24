package main

// 位运算
// 力扣：https://leetcode.cn/problems/the-two-sneaky-numbers-of-digitville/description/

// func getSneakyNumbers(nums []int) []int {
// 	// 注意：nums其实应该包含从 0 到 n - 1 的整数
// 	n := len(nums) - 2 // 设定 n 为原始数组元素的个数减去 2。这个前提是数组本来应该有 n 个唯一的值，但由于有两个重复的值，实际长度为 n + 2
// 	xorAll := 0        // 最终xorAll结果将会是两个重复数字之间的异或值（a ^ b）
// 	for i, x := range nums {
// 		xorAll ^= i ^ x
// 	}
// 	// 异或一下i没有到达的两个数字
// 	xorAll ^= n ^ (n + 1)
//
// 	// TrailingZeros 函数用于找出 xorAll 中最低位的 1 的索引。这个分界点将用来将数据分成两个组
// 	shift := bits.TrailingZeros(uint(xorAll))
//
// 	ans := make([]int, 2)
// 	for i, x := range nums {
// 		if i < n {
// 			// ans[i>>shift&1] ^= i：这个操作将当前索引 i 根据 shift 位进行右移并与 1 进行按位与（&）运算，从而决定将 i 放入 ans 数组的哪个位置。具体来说：
// 			// 如果 i 在这一位上是 0（i >> shift & 1 == 0），那么 i 会被异或到 ans[0]。
// 			// 如果 i 在这一位上是 1（i >> shift & 1 == 1），则会被异或到 ans[1]。
// 			ans[i>>shift&1] ^= i
// 		}
// 		ans[x>>shift&1] ^= x
// 	}
// 	return ans
// }

// 或者：
func getSneakyNumbers(nums []int) []int {
	total := 0
	for _, x := range nums {
		total ^= x
	}
	for i := 0; i <= len(nums)-1; i++ {
		total ^= i
	}
	// 异或一下 i 没有到达的两个数字
	total ^= (len(nums) - 2) ^ (len(nums) - 1)
	// 现在 total = a ^ b（a，b是两个结果值）

	total = total & (-total)
	a, b := 0, 0

	for i := 0; i < len(nums); i++ {
		if total&nums[i] == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	// 为什么 i < len(nums)-2 ？因为：数组本来的长度是 n 个，结果有两个元素重复出现了，所以长度变为了n+2。所以这里我们要 < len(nums)-2
	for i := 0; i < len(nums)-2; i++ {
		if total&i == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}
	return []int{a, b}
}
