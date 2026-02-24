package main

import (
	"slices"
)

// 二分查找：对答案进行二分查找
// 力扣：https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/description/

func smallestDivisor(nums []int, threshold int) int {
	// 结果值一定在 1 - nums[len(nums)-1] 之内
	slices.Sort(nums)

	res := 1

	l, r := 1, nums[len(nums)-1]

	for l <= r {
		mid := (l + r) >> 1

		// 判断mid作为除数是否可以
		total := 0
		for i := 0; i < len(nums); i++ {
			total += (nums[i] + mid - 1) / mid // 向上取整
		}

		if total <= threshold {
			r = mid - 1 // 要找最小的，所以继续往左边走
			res = mid
		} else if total > threshold {
			l = mid + 1 // 商的和太大了，我们要把除数调大一点，从而让商的和变小，所以往右边走
		}
	}
	return res
}

// func main() {
// 	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
// }
