package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/sliding-subarray-beauty/description/

// 超时
// func getSubarrayBeauty(nums []int, k int, x int) []int {
// 	result := make([]int, len(nums)-k+1)
// 	for right := 0; right < len(nums); right++ {
// 		if right >= k-1 {
// 			left := right - (k - 1)
// 			result[left] = getNumByX(nums[left:right+1], x)
// 		}
// 	}
//
// 	return result
// }
//
// // getNumByX 获取长度为k的滑动窗口排序后的美丽值
// func getNumByX(sli []int, x int) int {
// 	demo := make([]int, len(sli))
// 	copy(demo, sli)
// 	slices.Sort(demo)
// 	if demo[x-1] < 0 {
// 		return demo[x-1]
// 	}
// 	return 0
// }

// 优化：
// 注意题目范围：-50 <= nums[i] <= 50。
// 滑动窗口。由于值域很小，所以借鉴计数排序，用一个 cnt 数组维护窗口内每个数的出现次数。然后遍历 cnt 去求第 x 小的数。
// 什么是第 x 小的数？
// 设它是 num，那么 <num 的数有 <x 个，≤num 的数有 ≥x 个，就说明 num 是第 x 小的数。
func getSubarrayBeauty(nums []int, k int, x int) []int {
	result := make([]int, len(nums)-k+1)
	demo := [101]int{} // 存储的是 -50 <= nums[i] <= 50 出现了多少次

	for right := 0; right < len(nums); right++ {
		demo[nums[right]+50]++ // 把 索引-50 改为索引0

		if right >= k-1 {
			left := right - (k - 1)

			n := 0 // 统计符合x条件的元素
			for i := 0; i < len(demo); i++ {
				// fmt.Printf("i = %d, i-50 = %d, demo[i] = %d\n", i, i-50, demo[i])
				n += demo[i]
				if n >= x {
					if i-50 < 0 {
						// fmt.Println("------------------", i-50)
						result[left] = i - 50
						break
					}
				}
			}
			demo[nums[left]+50]--
		}
	}
	return result
}
