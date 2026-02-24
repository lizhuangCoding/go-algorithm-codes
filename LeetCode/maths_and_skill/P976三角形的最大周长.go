package main

import (
	"sort"
)

// 数学

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	maxRes := 0
	a, b, c := 0, 0, 0
	for i := 0; i <= len(nums)-3; i++ {
		a = nums[i]
		b = nums[i+1]
		c = nums[i+2]
		if a+b > c {
			maxRes = maxLargestPerimeter(maxRes, a+b+c)
		}
	}
	return maxRes
}

func maxLargestPerimeter(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
