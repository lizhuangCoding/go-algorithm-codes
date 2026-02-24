package main

// 位运算
// 力扣：https://leetcode.cn/problems/sum-of-all-subset-xor-totals/description/
// b站视频：https://www.bilibili.com/video/BV1YT4y117AH?spm_id_from=333.788.player.switch&vd_source=e2f6edb68f421b13075178535968e307&p=10

func subsetXORSum(nums []int) int {
	ans := 0 // 存储当前子集的异或值
	sum := 0 // 存储所有子集异或值的总和

	// 外层循环生成所有子集
	// 用左移运算符 << 计算出 2 的 len(nums) 次幂。因为一个包含 n 个元素的集合，其子集的数量为 2^n，所以 1 << len(nums) 代表子集的总数。
	// i 从 0 到 2^n - 1 进行遍历，每个 i 的二进制表示都能对应数组的一个子集。
	// 每次进入新的外层循环时，将 ans 重置为 0，以便计算新子集的异或值。
	for i := 0; i < (1 << len(nums)); i++ {
		ans = 0
		// 内层循环判断元素是否在当前子集内
		for j := 0; j < len(nums); j++ {
			// 对 i 和 1 << j 进行按位与运算。若结果不为 0，则表明 i 的二进制表示中第 j 位为 1，意味着数组中的第 j 个元素在当前子集中。
			if i&(1<<j) != 0 {
				// 若第 j 个元素在当前子集中，就将其与 ans 进行异或运算，更新 ans 的值。
				ans ^= nums[j]
			}
		}
		// 累加当前子集的异或值
		sum += ans
	}
	return sum
}
