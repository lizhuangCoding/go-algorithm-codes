package main

// 数学+数组
// 力扣：https://leetcode.cn/problems/add-to-array-form-of-integer/description/

func addToArrayForm(num []int, k int) []int {
	result := make([]int, 0)

	total := 0
	index := len(num) - 1

	for k > 0 || total > 0 || index >= 0 {
		demo := 0
		if index >= 0 {
			demo = num[index] + total + k%10
			index--
		} else {
			demo = total + k%10
		}
		k /= 10

		result = append(result, demo%10)
		total = demo / 10
	}

	// 反转切片
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
