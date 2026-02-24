package main

// 签到：数学

func minCount(coins []int) int {
	num := 0
	for _, v := range coins {
		num += (v + 1) / 2
	}
	return num
}
