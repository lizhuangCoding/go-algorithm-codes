package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/description/
//
// func fileCombination(target int) [][]int {
// 	res := make([][]int, 0)
//
// 	left, right := 1, 1
// 	sum := 1
//
// 	for left <= target/2 {
// 		if sum == target {
// 			demo := make([]int, 0)
// 			for i := left; i <= right; i++ {
// 				demo = append(demo, i)
// 			}
// 			res = append(res, demo)
//
// 			sum -= left
// 			left++
//
// 		} else if sum < target {
// 			right++
// 			sum += right
// 		} else {
// 			sum -= left
// 			left++
// 		}
// 	}
// 	return res
// }

// 或者：
func fileCombination(target int) [][]int {
	res := make([][]int, 0)

	left, right := 1, 2
	sum := 3

	for left < right {
		if sum == target {
			demo := make([]int, 0)
			for i := left; i <= right; i++ {
				demo = append(demo, i)
			}
			res = append(res, demo)

			sum -= left
			left++

		} else if sum < target {
			right++
			sum += right
		} else {
			sum -= left
			left++
		}
	}
	return res
}
