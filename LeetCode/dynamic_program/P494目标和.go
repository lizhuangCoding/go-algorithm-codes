package main

// 动态规划：01背包
// 力扣：https://leetcode.cn/problems/target-sum/description/

// 方法一：回溯（通过了，数组长度比较小）
// var result int
//
// func findTargetSumWays(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
//
// 	result = 0
//
// 	huisu(nums, target, 0, nums[0])
// 	huisu(nums, target, 0, -nums[0])
//
// 	return result
// }
//
// func huisu(nums []int, target, index, num int) {
// 	if index == len(nums)-1 {
// 		if num == target {
// 			result++
// 		}
// 		return
// 	}
//
// 	huisu(nums, target, index+1, num+nums[index+1])
// 	huisu(nums, target, index+1, num-nums[index+1])
// }

// 方法二：数学 + 01背包
// 把数组的数据分为两堆正数，一堆p1（加号），另一堆p2（减号）。所以 p1 + p2 = sum ，而且 p1 - p2 = target
// 所以合并两式，消去 p2：p1 + p2 + p1 - p2 = sum + target，得 p1 = (sum + target)/2
// 所以题目的要求可以变为：从数组中找到一堆数（和为p1）
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	s := (sum + target) / 2

	// 边界检查：sum + target 是奇数；sum < target
	if (sum+target)&1 == 1 || sum < findTargetSumWaysAbs(target) {
		return 0
	}

	// dp[i][j] 表示在前 i 个数中，和为 j 的值的个数
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, s+1)
	}

	// 初始化：前0个数凑出0的方案数为1
	dp[0][0] = 1

	// 如果当前数字 nums[i-1] 比目标和 j 大（j < nums[i-1]），则不能选这个数，组合数等于前 i-1 个数中和为 j 的组合数：
	// 如果当前数字 nums[i-1] 可以选（j >= nums[i-1]），则组合数等于：
	// 不选当前数的组合数：dp[i-1][j]；
	// 选当前数的组合数：dp[i-1][j-nums[i-1]]（即前 i-1 个数中和为 j - nums[i-1] 的组合数）。
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= s; j++ {

			// 容量不足，不选nums[i-1]
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				// 选或不选nums[i-1]的方案数之和
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}

		}
	}

	return dp[len(nums)][s]
}

func findTargetSumWaysAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// // 空间优化：一维
// func findTargetSumWays(nums []int, target int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
//
// 	// 边界条件检查
// 	if (target + total) % 2 != 0 || findTargetSumWaysAbs(target) > total {
// 		return 0
// 	}
//
// 	S := (target + total) / 2
//
// 	// dp[j] 表示凑出j的方案数
// 	dp := make([]int, S+1)
//
// 	// 初始化：凑出0的方案数为1（不选任何数）
// 	dp[0] = 1
//
// 	for _, num := range nums {
// 		// 必须倒序遍历，避免重复计算
// 		for j := S; j >= num; j-- {
// 			dp[j] += dp[j - num]
// 		}
// 	}
//
// 	return dp[S]
// }

// func main() {
// 	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)) // 5
// 	fmt.Println(findTargetSumWays([]int{1}, 1))             // 1
// 	fmt.Println(findTargetSumWays([]int{1}, 2))             // 0
// }
