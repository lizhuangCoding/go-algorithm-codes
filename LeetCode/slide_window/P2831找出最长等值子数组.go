package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/find-the-longest-equal-subarray/

// 时间复杂度：O(n)，其中 n 表示数组 nums 的长度。分组需要的时间复杂度为 O(n)，通过滑动窗口找到每种元素的最大长度只需遍历所有的连续相等字符的长度计数即可，最多有 n 个连续字符串的长度计数，因此总的时间复杂度为 O(n)。
// 空间复杂度：O(n)，其中 n 表示数组 nums 的长度。分组保存每种元素的索引序列需要的空间为 O(n)。
func longestEqualSubarray(nums []int, k int) int {
	pos := make(map[int][]int)
	for i, num := range nums {
		pos[num] = append(pos[num], i)
	}
	ans := 0
	for _, vec := range pos {
		left := 0
		for right := 0; right < len(vec); right++ {
			/* 缩小窗口，直到不同元素数量小于等于 k */
			// vec[right] - vec[left] +1：在nums数组中，vec[i]和vec[j]之间的距离（nums数组中，两个相同元素之间的距离）
			// right-left+1：vec[j]到vec[i]之间的元素个数（nums数组中，两个相同元素a之间，存在几个a，比如说 a，b，a，b，c，a，right-left+1的值为左边的a和右边的a的中间有几个a）
			// vec[i] - vec[j] - (i - j)：在 vec[j] 到 vec[i] 的索引范围内，有多少个不是目标元素的索引。
			for (vec[right]-vec[left]+1)-(right-left+1) > k {
				left++
			}
			if ans < right-left+1 {
				ans = right - left + 1
			}
		}
	}
	return ans
}
