package main

import (
	"strconv"
)

// 模拟
// 力扣：https://leetcode.cn/problems/baseball-game/description/

func calPoints(operations []string) int {
	scores := make([]int, 0)

	for _, v := range operations {
		if v == "C" {
			scores = scores[0 : len(scores)-1]
		} else if v == "D" {
			num := scores[len(scores)-1]
			scores = append(scores, 2*num)
		} else if v == "+" {
			num1, num2 := scores[len(scores)-2], scores[len(scores)-1]
			scores = append(scores, num1+num2)
		} else {
			num, _ := strconv.Atoi(v)
			scores = append(scores, num)
		}
	}

	res := 0
	for _, v := range scores {
		res += v
	}
	return res
}
