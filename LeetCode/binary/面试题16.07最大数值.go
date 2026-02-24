package main

// 位运算
// 力扣：https://leetcode.cn/problems/maximum-lcci/description/
// 题解：https://leetcode.cn/problems/maximum-lcci/solutions/279324/ji-jian-you-fu-hao-zheng-shu-wei-yi-fa-by-taichira/

func maximum(a int, b int) int {
	ret := int64(a - b)
	// ret&(ret>>63)的值：如果a>=b，那么为0，如果a<b，那么为a-b
	ret = int64(a) - ret&(ret>>63) // 根据符号位来选择 a 或 b
	return int(ret)                // 返回较大的数字
}

// 解释：
// ret >> 63 将 ret 右移 63 位，如果 ret 是负数（即 a < b），那么 ret >> 63 将得到 -1（即二进制全为 1）；如果 ret 是正数或零（即 a >= b），那么将得到 0。
// 接下来，ret & (ret >> 63) 会根据 ret 的符号位（正负）选择一个值：
// 如果 ret 是负数（a < b），那么 ret & -1 （即 ret）将被选择，意味着 a - b 的部分会被减去。（ret & -1 = ret）
// 如果 ret 是非负数（a >= b），ret & 0 将是 0，这时 int64(a) 不变。
