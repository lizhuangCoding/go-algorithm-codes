package main

// 签到：数组
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			res++
		}
	}
	return res
}
