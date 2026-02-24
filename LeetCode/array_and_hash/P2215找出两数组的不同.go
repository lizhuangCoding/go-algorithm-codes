package main

// 哈希，数组
// 力扣：https://leetcode.cn/problems/find-the-difference-of-two-arrays/

/*

错误写法：

问题1.重复元素处理不当
例子：nums1 = [1, 1, 2, 2]、nums2 = [2, 3]

你的代码：
m = {1:2, 2:2}，遍历 nums2: 遇到2时，m[2]-- 变为1
最终： result[0] = [1, 2] (但2其实在nums2中存在！)、result[1] = [3]

正确答案应该是：
result[0] = [1] (nums1有，nums2没有)
result[1] = [3] (nums2有，nums1没有)


问题2.缺少去重机制
第二个结果数组可能包含重复元素：
nums1 = [1]、nums2 = [2, 2, 2]
你的代码：result[1] = [2, 2, 2]
正确答案：result[1] = [2] (需要去重)

*/
// func findDifference(nums1 []int, nums2 []int) [][]int {
// 	m := make(map[int]int)
//
// 	for i := 0; i < len(nums1); i++ {
// 		m[nums1[i]]++
// 	}
//
// 	result := make([][]int, 2)
// 	result[0] = make([]int, 0)
// 	result[1] = make([]int, 0)
//
// 	for i := 0; i < len(nums2); i++ {
// 		demo := nums2[i]
// 		if _, ok := m[demo]; ok {
// 			m[demo]--
// 		} else {
// 			result[1] = append(result[1], demo)
// 		}
// 	}
//
// 	for k, v := range m {
// 		if v > 0 {
// 			result[0] = append(result[0], k)
// 		}
// 	}
//
// 	return result
// }

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	// 构建哈希集合
	for _, num := range nums1 {
		m1[num] = true
	}
	for _, num := range nums2 {
		m2[num] = true
	}

	res := make([][]int, 2)
	res[0] = make([]int, 0)
	res[1] = make([]int, 0)

	// 找到 nums1 中有但 nums2 中没有的元素
	for num := range m1 {
		if !m2[num] {
			res[0] = append(res[0], num)
		}
	}

	// 找到 nums2 中有但 nums1 中没有的元素
	for num := range m2 {
		if !m1[num] {
			res[1] = append(res[1], num)
		}
	}

	return res
}
