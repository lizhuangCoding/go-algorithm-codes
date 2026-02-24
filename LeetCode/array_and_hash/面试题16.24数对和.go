package main

// 哈希表
// 力扣：https://leetcode.cn/problems/pairs-with-sum-lcci/

// 时间超限
// func pairSums(nums []int, target int) [][]int {
// 	isVisited := make([]bool, len(nums))
//
// 	result := make([][]int, 0)
// 	for i := 0; i < len(nums); i++ {
//
// 		if !isVisited[i] {
// 			for j := i + 1; j < len(nums); j++ {
// 				if nums[i]+nums[j] == target && !isVisited[j] {
// 					result = append(result, []int{nums[i], nums[j]})
// 					isVisited[i] = true
// 					isVisited[j] = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// 优化：使用哈希表记录没有数字出现的次数
func pairSums(nums []int, target int) [][]int {
	var ans [][]int

	// 记录每个数字出现的次数
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	// 遍历 map 中的每个数字
	for num, count := range mp {
		// 计算当前数字与 target 的差值
		v := target - num

		// 判断差值是否存在于 map 中，并且当前数字与差值不相等
		if mp[v] > 0 && num != v {
			// 取当前数字和差值出现次数的最小值
			minD := pairSumsMin(count, mp[v])
			// 将当前数字和差值出现次数置为 0，避免重复计算
			mp[num] = 0
			mp[v] = 0
			// 将数对添加到结果中
			for i := 0; i < minD; i++ {
				ans = append(ans, []int{num, v})
			}

		} else if num == v && count >= 2 {
			// 如果当前数字与差值相等，并且出现次数大于等于 2
			for i := 0; i < count/2; i++ {
				ans = append(ans, []int{num, v})
			}
			// 将当前数字出现次数置为 0，避免重复计算
			mp[num] = 0
		}
	}
	return ans
}

func pairSumsMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
