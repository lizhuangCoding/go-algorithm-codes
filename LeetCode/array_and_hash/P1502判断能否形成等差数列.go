package main

import "sort"

// 签到：数组

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	cha := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != cha {
			return false
		}
	}
	return true
}
