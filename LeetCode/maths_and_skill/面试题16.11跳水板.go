package main

// 数学
// 力扣：https://leetcode.cn/problems/diving-board-lcci/description/

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{k * shorter}
	}

	res := make([]int, 0)

	n1, n2 := k, 0
	for n2 <= k {
		res = append(res, shorter*n1+longer*n2)
		n1--
		n2++
	}
	return res
}
