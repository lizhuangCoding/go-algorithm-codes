package main

import "sort"

// 签到：数组

func average(salary []int) float64 {
	sort.Ints(salary)
	sum := 0.0
	for i := 1; i <= len(salary)-2; i++ {
		sum += float64(salary[i])
	}
	return sum / float64(len(salary)-2)
}
