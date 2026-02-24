package cycle_or_recursion

// 签到：循环

func numIdenticalPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}
