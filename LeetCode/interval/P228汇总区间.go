package main

import (
	"strconv"
)

// 分组循环，数组
// 力扣：https://leetcode.cn/problems/summary-ranges/
// 灵茶山艾府 分组循环题单：https://leetcode.cn/problems/summary-ranges/solutions/553645/hui-zong-qu-jian-by-leetcode-solution-6zrs/comments/2106748

// func summaryRanges(nums []int) []string {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	result := make([]string, 0)
//
// 	num1, num2 := 0, 0
// 	num1 = nums[0]
//
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i+1] != nums[i]+1 {
// 			num2 = nums[i]
//
// 			if num1 != num2 {
// 				result = append(result, strconv.Itoa(num1)+"->"+strconv.Itoa(num2))
// 			} else {
// 				result = append(result, strconv.Itoa(num1))
// 			}
//
// 			num1 = nums[i+1]
// 		}
// 	}
//
// 	// 最后用一段是连续的情况
// 	num2 = nums[len(nums)-1]
// 	if num1 != num2 {
// 		result = append(result, strconv.Itoa(num1)+"->"+strconv.Itoa(num2))
// 	} else {
// 		result = append(result, strconv.Itoa(num1))
// 	}
//
// 	return result
// }

// 或者：
// 无需特判 nums 是否为空，也无需在循环结束后，再补上处理最后一段区间的逻辑。
func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; i++ {
		left := i

		for ; i < n-1 && nums[i]+1 == nums[i+1]; i++ {
		}

		s := strconv.Itoa(nums[left])
		if left < i {
			s += "->" + strconv.Itoa(nums[i])
		}
		ans = append(ans, s)
	}
	return
}
