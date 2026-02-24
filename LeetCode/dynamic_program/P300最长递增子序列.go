package main

// 动态规划
// 力扣：https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	result := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}

		}

		result = max(result, dp[i])
	}
	return result
}

// func main() {
// 	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
// 	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // 4
// 	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        // 1
// 	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})) // 6
// }
