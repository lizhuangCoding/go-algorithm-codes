package main

// 数组
// 力扣：https://leetcode.cn/problems/sort-colors/description/

// 方法一：单指针，遍历两次
// func sortColors(nums []int) {
// 	index := sortSortColors(nums, 0)
// 	sortSortColors(nums[index:], 1)
// }
//
// func sortSortColors(nums []int, v int) int {
// 	index := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == v {
// 			nums[index], nums[i] = nums[i], nums[index]
// 			index++
// 		}
// 	}
// 	return index
// }

// 方法二：双指针，遍历一次
// 题解：https://leetcode.cn/problems/sort-colors/solutions/437968/yan-se-fen-lei-by-leetcode-solution/
func sortColors(nums []int) {
	p0, p1 := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]

			// 重要（原因：看题解的图片）
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}

			p0++
			p1++

		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}

}
